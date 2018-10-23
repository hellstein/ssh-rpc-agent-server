package jobmgr

import (
//    "fmt"
    "log"
    "golang.org/x/crypto/ssh"
    "bytes"
    "strings"
)


type I_Job interface {
    Execute() (string, error)
    GetMachine() string
}


type Job struct {
    Machine I_Machine
    Tasks []I_Task
}

func (job *Job) GetMachine() string {
    return job.Machine.GetDomain()
}

func (job *Job) Execute() (string, error) {
    // Get client conf according to machine conf
    authConf, dest, err := job.Machine.GetAuthConf()
    if err != nil {
        return "", err
    }
    // Execute ssh session
    client, err := ssh.Dial("tcp", dest, authConf)
    if err != nil {
        log.Fatal("Failed to dial: ", err)
    }

    // Each ClientConn can support multiple interactive sessions,
    // represented by a Session.

    tasks := job.Tasks
    contents := []string{}
    for index, _ := range tasks {
        session, err := client.NewSession()
        if err != nil {
            log.Fatal("Failed to create session: ", err)
        }
        defer session.Close()

        // Once a Session is created, you can execute a single command on
        // the remote side using the Run method.
        var b bytes.Buffer
        session.Stdout = &b
        cmd := tasks[index].Serialize()
        if err := session.Run(cmd); err != nil {
            log.Fatal("Failed to run: " + err.Error())
        }
        contents = append(contents, b.String())
        b.Reset()
    }
    return strings.Join(contents[:],"\n"), nil
}
