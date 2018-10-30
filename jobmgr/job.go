package jobmgr

import (
//    "fmt"
    "log"
    "golang.org/x/crypto/ssh"
    "strings"
    "github.com/gorilla/websocket"
)


type I_Job interface {
    Execute(*websocket.Conn)
//    GetMachine() string
}


type Job struct {
    Machine I_Machine `json:machine`
    Tasks []I_Task `json:tasks`
}

func (job *Job) GetTaskCMD() string {
    ts := []string{}
    for index, _ := range job.Tasks {
        task := job.Tasks[index].Serialize()
        if strings.Contains(task, "sudo") {
            task = strings.Replace(task, "sudo", "MYPASS="+ job.Machine.GetSudoPassword() + " SUDO_ASKPASS=`echo $HOME`/echopass" +" sudo -A", -1)
            task = "echo -e \"#! /bin/bash\necho \\$MYPASS\">echopass && chmod +x echopass && " + task + " && rm `echo $HOME`/echopass"
        }
        ts = append(ts, task)
    }
    re :=strings.Join(ts, " && cd && ")
    log.Println("CMD: ", re)
    return re
}

/*
func (job *Job) GetMachine() string {
    return job.Machine.GetDomain()
}
*/

func (job *Job) Execute(conn *websocket.Conn) {
    // Get client conf according to machine conf
    authConf, dest, err := job.Machine.GetAuthConf()
    if err != nil {
       conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
    }
    // Execute ssh session
    job.RPC(dest, authConf, conn)
}

func syncIO (session *ssh.Session, conn *websocket.Conn, sshConn *ssh.Client) {
    go func(*ssh.Session, *websocket.Conn, *ssh.Client) {
        sessionReader, err := session.StdoutPipe()
        if err != nil {
          log.Fatal(err)
        }

        log.Println("======================== Sync session output ======================")
        defer func() {
            log.Println("======================== output: end ======================")
            conn.Close()
            sshConn.Close()
            session.Close()
        }()

        for {
            // set io.Writer of websocket
            outbuf := make([]byte, 4096)
            outn, err := sessionReader.Read(outbuf)
            if err != nil {
                log.Println("sshReader: ", err)
                return
            }
            err = conn.WriteMessage(websocket.TextMessage, outbuf[:outn])
            if err != nil {
                log.Println("connWriter: ", err)
                return
            }
        }
    }(session, conn, sshConn)

    go func(*ssh.Session, *websocket.Conn, *ssh.Client) {
        sessionWriter, err := session.StdinPipe()
        if err != nil {
            log.Fatal(err)
        }

        log.Println("======================== Sync session input ======================")
        defer func() {
            log.Println("======================== input: end ======================")
            conn.Close()
            sshConn.Close()
            session.Close()
        }()

        for {
            // set up io.Reader of websocket
            _, reader, err := conn.NextReader()
            if err != nil {
                log.Println("connReaderCreator: ", err)
                return
            }
            buf := make([]byte, 1024)
            n, err := reader.Read(buf)
            if err != nil {
                log.Print("connReader: ", err)
                return
            }
            _, err = sessionWriter.Write(buf[:n])
            if err != nil {
                log.Print("sshWriter: ", err)
                conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
                return
            }
        }
    }(session, conn, sshConn)
}

func (job *Job) RPC(dest string, authConf *ssh.ClientConfig, conn *websocket.Conn) {
    sshConn, err := ssh.Dial("tcp", dest, authConf)
    if err != nil {
        log.Fatal("Failed to dial: ", err)
    }
    // Once a Session is created, you can execute a single command on
    // the remote side using the Run method.
    // Set New session
    ss, err := sshConn.NewSession()
    if err != nil {
        log.Fatal("Failed to create session: ", err)
    }
    modes := ssh.TerminalModes {
      ssh.ECHO:          0,     // disable echoing
      ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
      ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    // Request pseudo terminal
    if err := ss.RequestPty("xterm", 40, 80, modes); err != nil {
        log.Fatal("request for pseudo terminal failed: ", err)
    }

    syncIO(ss, conn, sshConn)
    cmd := job.GetTaskCMD()
    if err := ss.Run(cmd); err != nil {
        log.Fatal("failed to run cmd: ", err)
    }
}


