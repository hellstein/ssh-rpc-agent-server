package jobmgr

import (
    "log"
    "golang.org/x/crypto/ssh"
    "strings"
//    "golang.org/x/crypto/ssh/terminal"
//    "os"
    "github.com/gorilla/websocket"
)

/*
    Define interface I_Job with only method Execute
*/
type I_Job interface {
    Execute(*websocket.Conn)
}


/*
    Implement I_Job
*/
type Job struct {
    Machine I_Machine
    Tasks []I_Task
}

/*
    Formalize cmd of tasks for ssh
*/
func (job *Job) GetTaskCMD() string {
    ts := []string{}
    for index, _ := range job.Tasks {
        task := job.Tasks[index].Serialize()
        if strings.Contains(task, "sudo") && strings.Index(task, "sudo")==0 {
            task = strings.Replace(task, "sudo", "MYPASS="+ job.Machine.GetSudoPassword() + " SUDO_ASKPASS=`echo $HOME`/echopass" +" sudo -A", 1)
            task = "echo -e \"#! /bin/bash\necho \\$MYPASS\">echopass && chmod +x echopass && " + task + " && rm `echo $HOME`/echopass"
        }
        ts = append(ts, task)
    }
    re := strings.Join(ts, " && cd && ")
    log.Println("CMD: ", re)
    return re
}


/*
    Get ssh session and ssh client
*/
func (job *Job) GetSSH() (*ssh.Session, *ssh.Client) {

    // Get client conf according to machine conf
    authConf, dest, err := job.Machine.GetAuthConf()
    if err != nil {
       log.Println(err)
    }
    client, err := ssh.Dial("tcp", dest, authConf)
    if err != nil {
        log.Fatal("Failed to dial: ", err)
    }

    // Once a Session is created, you can execute a single command on
    // the remote side using the Run method.
    // Set New session
    session, err := client.NewSession()
    if err != nil {
        log.Fatal("Failed to create session: ", err)
    }

    modes := ssh.TerminalModes {
      ssh.ECHO:          1,     // disable echoing
      ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
      ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    err = session.RequestPty("xterm-256color", 55, 110, modes)
    if err != nil {
        log.Fatal("request for pseudo terminal failed: ", err)
    }

    return session, client
}

/*
    Execute ssh job and communicate through websocket
*/
func (job *Job) Execute(conn *websocket.Conn) {

    session, client := job.GetSSH()

    defer client.Close()
    defer session.Close()
    syncIO(session, client, conn)

    // Start remote shell
    cmd := job.GetTaskCMD()

    if err := session.Run(cmd); err != nil {
        log.Println("failed to start shell: ", err)
        return
    }
}

