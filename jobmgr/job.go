package jobmgr

import (
//    "fmt"
    "log"
    "golang.org/x/crypto/ssh"
    "bytes"
    "strings"
    "encoding/json"
//    "os"
)


type I_Job interface {
    Execute(chan string)
    GetMachine() string
}


type Job struct {
    Machine I_Machine
    Tasks []I_Task
}

func (job *Job) GetTasks() []string {
    ts := []string{}
    for index, _ := range job.Tasks {
        task := job.Tasks[index].Serialize()
        if strings.Contains(task, "sudo") {
            task = strings.Replace(task, "sudo", "MYPASS="+ job.Machine.GetSudoPassword() + " SUDO_ASKPASS=`echo $HOME`/echopass" +" sudo -A", -1)
            task = "echo -e \"#! /bin/bash\necho \\$MYPASS\">echopass && chmod +x echopass && " + task + " && rm `echo $HOME`/echopass"
        }
        ts = append(ts, task)
    }
    return ts
}

func (job *Job) GetMachine() string {
    return job.Machine.GetDomain()
}

func (job *Job) Execute(result chan string) {
    // Get client conf according to machine conf
    authConf, dest, err := job.Machine.GetAuthConf()
    if err != nil {
       result <- err.Error()
    }
    // Execute ssh session
    job.RPC(dest, authConf, result)
}

func (job *Job) RPC(dest string, authConf *ssh.ClientConfig, result chan string) {
    client, err := ssh.Dial("tcp", dest, authConf)
    if err != nil {
        log.Fatal("Failed to dial: ", err)
    }

    // Each ClientConn can support multiple interactive sessions,
    // represented by a Session.
    tasks := job.Tasks
    content := Result{Machine: job.GetMachine(), TRS: []TaskResult{}}
    for index, _ := range tasks {
        // Once a Session is created, you can execute a single command on
        // the remote side using the Run method.
        session, err := client.NewSession()
        if err != nil {
            log.Fatal("Failed to create session: ", err)
        }
        defer session.Close()

        modes := ssh.TerminalModes{
          ssh.ECHO:          0,     // disable echoing
          ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
          ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
        }
    // Request pseudo terminal
        if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
            log.Fatal("request for pseudo terminal failed: ", err)
        }
        tr := TaskResult{Topic: tasks[index].GetTopic()}
        var b bytes.Buffer
        session.Stdout = &b
        cmd := job.GetTasks()[index]
        if err := session.Run(cmd); err != nil {
            tr.Msg = err.Error()
        } else {
            tr.Msg = b.String()
        }
        content.TRS = append(content.TRS, tr)
        b.Reset()
    }
    re, _ := json.Marshal(&content)
    result <- string(re)
}


