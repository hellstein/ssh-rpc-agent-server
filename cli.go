package main

import (
//    "fmt"
    "github.com/urfave/cli"
    "os"
    "log"
)


/*
    Create flags
*/
func InitFlags() []cli.Flag {

    return  []cli.Flag {
        cli.StringFlag {
            Name: "machinefile, mf",
            Usage: "Specify the machine configuration file",
        },
        cli.StringFlag {
            Name: "taskfile, tf",
            Usage: "Specify the task configuration file",
        },
    }

}

/*
    Handler argument validation
*/
func argsHandler(c *cli.Context) error {
    if len(os.Args)==1 {
        return cli.NewExitError("Error: No flags specified.\nPlease use --tf to specify tasks and --mf to specify machines", 40)
    }
    if c.String("machinefile") == "" {
        return cli.NewExitError("Error: No configuration file of machines has been specified.\nPlease use --mf to specify machines", 41)
    }
    if c.String("taskfile") == "" {
        return cli.NewExitError("Error: No configuration file of tasks has been specified.\nPlease use --tf to specify tasks", 42)
    }
    return nil
}

/*
    Generate a function for app.Action usage
*/
func InitAppAction(argsHandler func(*cli.Context) error, invoker *Invoker) func(*cli.Context) error {
    return func(c *cli.Context) error {
        if err := argsHandler(c); err!=nil {
            return err
        }
        invoker.invoke(c.String("taskfile"), c.String("machinefile"))
        return nil
    }
}


/*
    Initialize cli management
*/
func InitCli() {
    app := cli.NewApp()
    invoker := &Invoker{}

    app.Usage = "RPC support tool"
    app.Version = "0.0.2"
    app.Flags = InitFlags()
    app.Action = InitAppAction(argsHandler, invoker)

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }

}
