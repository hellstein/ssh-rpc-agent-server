package jobmgr

import (
//    "encoding/json"
)

type I_Result interface {
}

type TaskResult struct {
    Topic string `json:"topic"`
    Msg string `json:"msg"`
}

type Result struct {
    Machine string `json:"machine"`
    TRS []TaskResult `json:"taskresults"`
}
