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
        results := mgr.ExecuteJobs(jobs)
        // Output results
        fmt.Fprint(w, results)
    }
}

