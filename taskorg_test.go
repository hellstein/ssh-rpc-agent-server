package main

import (
    "testing"
    "github.com/hellstein/ssh-rpc-agent/assignment"

)


type MockTask struct {
}

/*
    Implement I_Task methods

*/
func (m *MockTask) Serialize() string {
    return ""
}

func (m *MockTask) GetSummary() string {
    return ""
}



func TestGetTaskSet(t *testing.T) {
    cases := []struct{
        name string
        to I_TaskOrg
        eLen int
    }{
        {name: "no tasks", to: &TaskOrg{}, eLen:0},
        {name: "one task", to: &TaskOrg{Tset: []assignment.I_Task{&MockTask{}}}, eLen:1},
        {name: "more tasks", to: &TaskOrg{Tset: []assignment.I_Task{&MockTask{},&MockTask{}}}, eLen:2},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            tLen := len(tc.to.GetTaskSet())
            if tLen != tc.eLen {
                t.Fatalf("expected number of tasks is %v, but got %v", tc.eLen, tLen)
            }
        })
    }

}


func TestParseTask(t *testing.T) {
    cases := []struct{
        name string
        file string
        eLen int
    }{
        {name: "no file", file: "", eLen:0},
        {name: "with file which doesn't exist ", file: "xxx", eLen:0},
        {name: "with correct file", file: "tasks_test.json", eLen:2},
    }

    to := &TaskOrg{}
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            tLen := len(to.parse(tc.file))
            if tLen != tc.eLen {
                t.Fatalf("expected number of machines is %v, but got %v", tc.eLen, tLen)
            }
        })
    }


}



