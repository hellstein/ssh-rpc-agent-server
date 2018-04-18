package main


import (
    "github.com/dorrywhale/customer-support/assignment"
)


/*
    Define interface: I_MachineOrg
*/
type I_MachineOrg interface {
    GetMachineSet() []assignment.I_Machine
}


/*
    Define class: MachineOrg Implementing I_MachineOrg
*/
type MachineOrg struct {
    Mset []assignment.I_Machine
}


/*
   Implement I_MachineOrg.GetMachineSet() []assignment.I_Machine
*/
func (morg *MachineOrg) GetMachineSet() []assignment.I_Machine {
    return morg.Mset
}


/*
    Implement I_MachineOrg.ParseMachineFile(filename string) []assignment.I_Machine
*/
func (morg *MachineOrg) parse(filename string) []assignment.I_Machine {
    return assignment.ParseMachineFile(filename)
}

