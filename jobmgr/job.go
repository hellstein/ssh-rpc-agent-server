package jobmgr

import (
//    "fmt"
    "log"
    "golang.org/x/crypto/ssh"
    "bytes"
    "strings"
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
    content := []string{job.GetMachine(),}
    for index, _ := range tasks {
        // Once a Session is created, you can execute a single command on
        // the remote side using the Run method.
        session, err := client.NewSession()
        if err != nil {
            log.Fatal("Failed to create session: ", err)
        }
        defer session.Close()

        var b bytes.Buffer
        session.Stdout = &b
        content = append(content, tasks[index].GetTopic())
        cmd := job.GetTasks()[index]
        if err := session.Run(cmd); err != nil {
            content = append(content, err.Error())
        } else {
            content = append(content, b.String())
        }
        b.Reset()
    }
    result <- strings.Join(content, "\n")
}


