package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/hellstein/ssh-rpc-agent/jobmgr"
)




func getHandler(mgr jobmgr.I_Mgr) func(http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        // Get machine conf and task conf from request
        conf, err := GetConf(r)
        if err != nil {
           log.Fatal(err)
        }
        mconf := conf["machines"]
        tconf := conf["tasks"]
        // Create job for each machine
        jobs := mgr.CreateJobs(mconf, tconf)
        // Execute jobs
        result := make(chan string, len(jobs))
        mgr.ExecuteJobs(jobs, result)
        cter := 0
        for i := range result {
            fmt.Fprint(w, i)
            cter = cter + 1
            if cter == len(jobs) {
                close(result)
            }
        }
    }
}

