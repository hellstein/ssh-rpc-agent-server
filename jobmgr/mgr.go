package jobmgr

import (
//    "fmt"
)


type I_Mgr interface {
    CreateJobs([]byte, []byte) []I_Job
    ExecuteJobs([]I_Job) map[string]string
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

func (mgr *Mgr) ExecuteJobs(jobs []I_Job) map[string]string {
    results := map[string]string{}
    var err error
    var msg string
    for index, _ := range jobs {
        msg, err = jobs[index].Execute()
        if err != nil {
            msg = err.Error()
        }
        results[jobs[index].GetMachine()] = msg
    }
    return results
}
