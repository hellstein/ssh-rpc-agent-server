package command


import (
    "testing"
    "reflect"
    "os/exec"
)

var e I_Executor
var cmd I_Command

type MockCommand struct {
    Presentation []string
    Title string
}

func (c *MockCommand) Execute(e I_Executor) error {
    return nil
}

func (c *MockCommand) GetPresentation() []string {
    return c.Presentation
}



func (c *MockCommand) GetTitle() string {
    return c.Title
}

func setupExecutorTestCase(t *testing.T) func(*testing.T) {
    e = &Executor{}
    t.Log("setup test case")
    return func(t *testing.T) {
        e = nil
        t.Log("teardown test case")
    }
}

func setupExecutorSubTest(t *testing.T, c I_Command) func(*testing.T) {
    cmd = c
    t.Log("setup sub test")
    return func(t *testing.T) {
        cmd = nil
        t.Log("teardown sub test")
    }
}

func TestExecutorExecute(t *testing.T) {
    cases := []struct {
        name string
        c I_Command
        result error
    }{
        {
            name: "with correct configuration",
            c: &MockCommand{
                Presentation: []string{"ssh", "-p", "22", "-o", "StrictHostKeyChecking=no", "-t", "root@127.0.0.1", "/bin/bash", "-c", "ls"},
                Title: "list files",
            },
            result:nil,
        },
        {
            name: "without task",
            c: &MockCommand{
                Presentation: []string{"ssh", "-p", "22", "-o", "StrictHostKeyChecking=no", "-t", "root@127.0.0.1", "/bin/bash", "-c"},
                Title: "list files",
            },
            result: &exec.ExitError{},
        },
        {
            name: "without machine",
            c: &MockCommand{
                Presentation: []string{"ssh", "-t", "/bin/bash", "-c"},
                Title: "no task",
            },
            result: &exec.ExitError{},
        },

    }
    teardownExecutorTestCase := setupExecutorTestCase(t)
    defer teardownExecutorTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownExecutorSubTest := setupExecutorSubTest(t, tc.c)
            defer teardownExecutorSubTest(t)
            err := e.Execute(cmd)
            if reflect.TypeOf(err) != reflect.TypeOf(tc.result) {
                t.Fatalf("expected result %v, but got %v", tc.result, err)
            }
        })
    }
}

