package main

import (
    "testing"
    "reflect"
    "github.com/hellstein/ssh-rpc-agent/assignment"
)




func TestGenCMDs(t *testing.T) {
    cases := []struct{
        name string
        tasks []assignment.I_Task
        machines []assignment.I_Machine
        eLen int
    }{
        {"with no tasks or machines", nil, nil, 0},
        {"with empty tasks and no machines", []assignment.I_Task{}, nil, 0},
        {"with no tasks and empty machines", nil, []assignment.I_Machine{}, 0},
        {"with empty tasks and machines", []assignment.I_Task{}, []assignment.I_Machine{}, 0},
        {
            "with more tasks and machines",
            []assignment.I_Task{&MockTask{}, &MockTask{}},
            []assignment.I_Machine{&MockMachine{}, &MockMachine{}, &MockMachine{}},
            6,
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            cLen := len(genCMDs(tc.tasks, tc.machines))
            if cLen != tc.eLen {
                t.Fatalf("expected number of commands is %v, but got %v", tc.eLen, cLen)
            }
        })
    }
}

// TODO
func TestFormalizeCMDs(t *testing.T) {
}


// TODO

func TestRun(t *testing.T) {
}


func TestInitMachineOrg(t *testing.T) {
    cases := []struct{
        name string
        file string
        eLen int
    }{
        {name: "no file", file: "", eLen:0},
        {name: "with file which doesn't exist ", file: "xxx", eLen:0},
        {name: "with correct file", file: "machines_test.json", eLen:3},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            mLen := len(InitMachineOrg(tc.file).GetMachineSet())
            if mLen != tc.eLen {
                t.Fatalf("expected number of machines is %v, but got %v", tc.eLen, mLen)
            }
        })
    }

}

func TestInitTaskOrg(t *testing.T) {
    cases := []struct{
        name string
        file string
        eLen int
    }{
        {name: "no file", file: "", eLen:0},
        {name: "with file which doesn't exist ", file: "xxx", eLen:0},
        {name: "with correct file", file: "tasks_test.json", eLen:2},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            tLen := len(InitTaskOrg(tc.file).GetTaskSet())
            if tLen != tc.eLen {
                t.Fatalf("expected number of machines is %v, but got %v", tc.eLen, tLen)
            }
        })
    }


}


func TestInitExecutor(t *testing.T) {
    e := InitExecutor()
    eType := reflect.TypeOf(e).Kind()
    if  eType != reflect.Ptr {
        t.Errorf("The object type of InitExecutor() should be %v, but got %v", reflect.Ptr, eType)
    }
}



