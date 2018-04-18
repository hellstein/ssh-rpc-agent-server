package command


import (
   "fmt"
    "os"
    "os/exec"
)

/*
    Define interface: I_Executor
*/
type I_Executor interface {
    Execute(cmd I_Command) error
}


/*
    Define class: Executor implementing I_Executor
*/
type Executor struct {
}


/*
    Execute I_Command
*/
func (e *Executor) Execute(cmd I_Command) error {
    prog := cmd.GetPresentation()[0]
    args := cmd.GetPresentation()[1:]

    exe := exec.Command(prog, args...)
    exe.Stdout = os.Stdout
    exe.Stderr = os.Stderr
    exe.Stdin = os.Stdin
    fmt.Println(cmd.GetTitle())
    err := exe.Run()
    if err != nil {
        fmt.Printf("Failed to start %v.\n Error: \n %s\n", exe, err.Error())
    }

    return err
}


