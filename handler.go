package main

import (
//    "fmt"
    "net/http"
    "log"
    "github.com/hellstein/ssh-rpc-agent/jobmgr"
    "github.com/gorilla/websocket"
)




func getHandler(mgr jobmgr.I_Mgr) func(http.ResponseWriter, *http.Request) {
    var upgrader = websocket.Upgrader{}

    return func (w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
          log.Println("Upgrade:", err)
          return
        }

        mgr.SetConn(conn)
        mgr.ExecuteJob()
    }
}

