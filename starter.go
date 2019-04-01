package main
import (
    "log"
    "net/http"
    "github.com/hellstein/ssh-rpc-agent-server/jobmgr"
)

func main() {
    mgr := &jobmgr.Mgr{}
    log.Println("Serving Agent on 0.0.0.0 port 8900 ...")
    log.Fatal(http.ListenAndServe(":8900", CreateRouter(mgr)))
}
