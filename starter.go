package main
import (
    "log"
    "net/http"
    "github.com/hellstein/ssh-rpc-agent/jobmgr"
)

func main() {
    mgr := &jobmgr.Mgr{}
    log.Fatal(http.ListenAndServe(":8900", CreateRouter(mgr)))
}
