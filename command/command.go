package command


import (
    "fmt"
    "github.com/dorrywhale/customer-support/assignment"
    "strings"
)


/*
    Define interface: I_Command
*/
type I_Command interface {
    Execute(e I_Executor) error
    GetPresentation() []string
    GetTitle() string
}


/*
    Define class: Command implementing I_Command
*/
type Command struct {
    Machine assignment.I_Machine
    Task assignment.I_Task
}


/*
    Execute command by I_Executor
*/
func (c *Command) Execute(e I_Executor) error {
    return e.Execute(c)
}


/*
    Get presentation of machine according to its mode
*/
func (c *Command) GetPresentation() []string {
    re := []string{}
    if c.Machine.GetMode() == "SSHKEY" {
        re = append(re, "ssh", "-t")
        re = append(re, c.Machine.GetLabel())
    }
    if c.Machine.GetMode() == "USERPASS" {
        re = append(re, "sshpass", "-p", c.Machine.GetSudoPassword(), "ssh", "-t")
        re = append(re, "-p", c.Machine.GetPort())
        re = append(re, c.Machine.GetUsername() + "@" + c.Machine.GetDomain())
    }
    re = append(re, "/bin/bash", "-c", c.GetTask())
    return re
}

/*
    Get task
*/
func (c *Command) GetTask() string {
    task := strings.Replace(c.Task.Serialize(), "sudo", "echo "+c.Machine.GetSudoPassword()+" | sudo -S", -1)
    return "'"+task+"'"
}

/*
    Get task title
*/
func (c *Command) GetTitle() string {
    return fmt.Sprintf(strings.Repeat("=", 40)+"\n%s\n"+strings.Repeat("=", 40)+"\n",  c.Task.GetSummary())
}
