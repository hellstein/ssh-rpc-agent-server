package command


import (
    "fmt"
    "github.com/FuQiFeiPian/ssh-rpc-agent/assignment"
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
        re = append(re, "ssh", "-o", "StrictHostKeyChecking=no", "-t")
        re = append(re, c.Machine.GetLabel())
    }
    if c.Machine.GetMode() == "USERPASS" {
        re = append(re, "sshpass", "-p", c.Machine.GetSudoPassword(), "ssh", "-o", "StrictHostKeyChecking=no", "-t")
        re = append(re, "-p", c.Machine.GetPort())
        re = append(re, c.Machine.GetUsername() + "@" + c.Machine.GetDomain())
    }
    re = append(re, "/bin/bash", "-l", "-c", c.GetTask())
    return re
}

/*
    Get task
*/
func (c *Command) GetTask() string {
    task := c.Task.Serialize()
    if strings.Contains(task, "sudo") {
        task = strings.Replace(task, "sudo", "MYPASS="+c.Machine.GetSudoPassword() + " SUDO_ASKPASS=`echo $HOME`/echopass" +" sudo -A", -1)
        task = "echo -e \"#! /bin/bash\necho \\$MYPASS\">echopass && chmod +x echopass && " + task + " && rm `echo $HOME`/echopass"
    }
    return "'"+task+"'"
}

/*
    Get task title
*/
func (c *Command) GetTitle() string {
    return fmt.Sprintf(strings.Repeat("=", 40)+"\n%s\n"+strings.Repeat("=", 40)+"\n",  c.Task.GetSummary())
}
