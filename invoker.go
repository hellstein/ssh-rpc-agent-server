package main



import (

)



type Invoker struct {
}


/*
    Initialize the processor by task and machine configuration file
*/
func InitProcessor(tfile string, mfile string) I_Processor {
    return &Processor{
        executor: InitExecutor(),
        torg: InitTaskOrg(tfile),
        morg: InitMachineOrg(mfile),
    }
}

/*
    Invoker entry point
*/
func (in *Invoker) invoke(tfile string, mfile string) []error {
    proc := InitProcessor(tfile, mfile)
    return proc.run(proc.formalizeCMDs())
}
