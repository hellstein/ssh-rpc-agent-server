package main


import (
    "github.com/dorrywhale/ssh-rpc-agent/assignment"
)


/*
    Define interface: I_TaskOrg
*/
type I_TaskOrg interface {
    GetTaskSet() []assignment.I_Task
}

/*
    Define class: TaskOrg Implementing I_TaskOrg
*/
type TaskOrg struct {
    Tset []assignment.I_Task
}

/*
   Implement I_MachineOrg.GetMachineSet(filter map[string]string) []assignment.I_Machine
*/
func (torg *TaskOrg) GetTaskSet() []assignment.I_Task {
    return torg.Tset
}

/*
    Implement I_TaskOrg.ParseTaskFile(filename string) []assignment.I_Task
*/
func (torg *TaskOrg) parse(filename string) []assignment.I_Task {
    return assignment.ParseTaskFile(filename)
}


