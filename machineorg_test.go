package main

import (
    "testing"
    "github.com/hellstein/ssh-rpc-agent/assignment"

)


type MockMachine struct {
}

/*
    Implement I_Machine methods

*/
func (m *MockMachine) GetLabel() string {
    return ""
}

func (m *MockMachine) GetMode() string {
    return ""
}

func (m *MockMachine) GetSudoPassword() string {
    return ""
}

func (m *MockMachine) GetUsername() string {
    return ""
}

func (m *MockMachine) GetDomain() string {
    return ""
}


func (m *MockMachine) GetPort() string {
    return ""
}


func TestGetMachineSet(t *testing.T) {
    cases := []struct{
        name string
        mo I_MachineOrg
        eLen int
    }{
        {name: "no machines", mo: &MachineOrg{}, eLen:0},
        {name: "one machine", mo: &MachineOrg{Mset: []assignment.I_Machine{&MockMachine{}}}, eLen:1},
        {name: "more machines", mo: &MachineOrg{Mset: []assignment.I_Machine{&MockMachine{},&MockMachine{}}}, eLen:2},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            mLen := len(tc.mo.GetMachineSet())
            if mLen != tc.eLen {
                t.Fatalf("expected number of machines is %v, but got %v", tc.eLen, mLen)
            }
        })
    }
}

func TestParseMachine(t *testing.T) {
    cases := []struct{
        name string
        file string
        eLen int
    }{
        {name: "no file", file: "", eLen:0},
        {name: "with file which doesn't exist ", file: "xxx", eLen:0},
        {name: "with correct file", file: "machines_test.json", eLen:3},
    }

    mo := &MachineOrg{}
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            mLen := len(mo.parse(tc.file))
            if mLen != tc.eLen {
                t.Fatalf("expected number of machines is %v, but got %v", tc.eLen, mLen)
            }
        })
    }

}



