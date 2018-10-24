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
        mconf := conf["machine"]
        tconf := conf["tasks"]
        // Create job for each machine
        jobs := mgr.CreateJobs(mconf, tconf)
        // Execute jobs
        result := make(chan string)
        mgr.ExecuteJobs(jobs, result)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, <-result)
    }
}

