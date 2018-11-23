package main
import (
//    "log"
//    "net/http"
    "github.com/hellstein/ssh-rpc-agent/jobmgr"
)

func main() {
    mgr := &jobmgr.Mgr{}

    machine := jobmgr.Machine {
        Domain: "192.168.21.5",
        Port: "22",
        Username: "develop_pan",
        Mode: "USERPASS",
        Credential: jobmgr.Credential{
            Password: "abc123_",
        },
        SudoPassword: "abc123_",
    }
    job := &jobmgr.Job{Machine: &machine, Tasks: []jobmgr.I_Task{}}
    task := jobmgr.Task {
        Topic: "Test",
        Tasks: []string{"sudo iftop"},
    }
    job.Tasks = append(job.Tasks, &task)
    mgr.Job = job


    mgr.ExecuteJob()
 //   log.Fatal(http.ListenAndServe(":8900", CreateRouter(mgr)))
}
