package jobmgr

import (
//    "fmt"
//    "net/http"
)


type I_Mgr interface {
    CreateJobs([]byte, []byte) []I_Job
    ExecuteJobs([]I_Job, chan string)
}


type Mgr struct {
}

func (mgr *Mgr) CreateJobs(mconf []byte, tconf []byte) []I_Job {
    // machine interface
    machines := ParseMachineData(mconf)
    // task interface
    tasks := ParseTaskData(tconf)
    // create jobs
    jobs := []I_Job{}
    for index, _ := range machines {
        jobs = append(jobs, &Job{Machine: machines[index], Tasks: tasks})
    }
    return jobs
}

func (mgr *Mgr) ExecuteJobs(jobs []I_Job, result chan string) {
    for index, _ := range jobs {
        //machine := jobs[index].GetMachine()
        go jobs[index].Execute(result)
    }
}
