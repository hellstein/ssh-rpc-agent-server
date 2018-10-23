package jobmgr

import (
//    "fmt"
//    "net/http"
)


type I_Mgr interface {
    CreateJobs([]byte, []byte) I_Job
    ExecuteJobs(I_Job, chan string)
}


type Mgr struct {
}

func (mgr *Mgr) CreateJobs(mconf []byte, tconf []byte) I_Job {
    // machine interface
    machine := ParseMachineData(mconf)
    // task interface
    tasks := ParseTaskData(tconf)
    // create jobs
    job := &Job{Machine: machine, Tasks: tasks}
    return job
}

func (mgr *Mgr) ExecuteJobs(job I_Job, result chan string) {
    go job.Execute(result)
}
