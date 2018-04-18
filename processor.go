package main


import (
    "github.com/dorrywhale/customer-support/command"
    "github.com/dorrywhale/customer-support/assignment"
    "fmt"
)


/*
    Define interface: I_Processor
*/
type I_Processor interface {
    formalizeCMDs() []command.I_Command
    run(cmds []command.I_Command) []error
}

/*
    Define class: Processor implementing I_Processor
*/
type Processor struct {
    executor command.I_Executor
    torg I_TaskOrg
    morg I_MachineOrg
}


/*
    Generate []I_Command by []I_Task and []I_Machine
*/
func genCMDs(tasks []assignment.I_Task, machines []assignment.I_Machine) []command.I_Command {
    fmt.Printf("You have specify %d tasks on %d machines\n", len(tasks), len(machines))
    cmds := []command.I_Command{}

    for _, m := range machines {
        for _, t := range tasks {
            cmds = append(cmds, &command.Command{Task: t, Machine: m})
        }
    }
    return cmds
}

/*
    Formalize the []I_Command 
*/
func (proc *Processor) formalizeCMDs() []command.I_Command {
    tasks := proc.torg.GetTaskSet()
    machines := proc.morg.GetMachineSet()
    return genCMDs(tasks, machines)
}


/*
    Execute the []I_Command
*/
func (proc *Processor) run(cmds []command.I_Command) []error {
    result := []error{}
    for _, cmd := range cmds {
        result = append(result, cmd.Execute(proc.executor))
    }
    return result
}

/*
    Initialize the MachineOrg
*/
func InitMachineOrg(mfile string) I_MachineOrg {
    morg := &MachineOrg{Mset: []assignment.I_Machine{}}
    if mfile !="" {
        morg.Mset = morg.parse(mfile)
    }
    return morg
}


/*
    Initialize the TaskOrg
*/

func InitTaskOrg(tfile string) I_TaskOrg {
    torg := &TaskOrg{Tset: []assignment.I_Task{}}
    if tfile !="" {
        torg.Tset = torg.parse(tfile)
    }

    return torg
}

/*
    Initialize the I_Executor
*/
func InitExecutor() command.I_Executor {
    return &command.Executor{}
}
