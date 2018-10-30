package main

import (
    "github.com/gorilla/mux"
    "github.com/hellstein/ssh-rpc-agent/jobmgr"
)

func CreateRouter(mgr jobmgr.I_Mgr) *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/test", getHandler(mgr))
    return router
}
