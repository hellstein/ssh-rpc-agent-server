package command


import (
    "github.com/dorrywhale/ssh-rpc-agent/assignment"
    "testing"
    "os/exec"
    "reflect"
    "strings"
)


type MockExecutorSucc struct {
}

func (me *MockExecutorSucc) Execute(c I_Command) error {
    return nil
}

type MockExecutorFail struct {
}

func (me *MockExecutorFail) Execute(c I_Command) error {
    return &exec.ExitError{}
}

func setupCommandTestCase(t *testing.T) func(t *testing.T) {
    t.Log("setup test case")
    return func(t *testing.T) {
        t.Log("teardown test case")
    }
}

func setupCommandSubTest(t *testing.T) func(t *testing.T) {
    t.Log("setup sub test")
    return func(t *testing.T) {
        t.Log("teardown sub test")
    }
}



type MockMachine struct {
    Label string
    Domain string
    Username string
    Port string
    SudoPassword string
    Mode string
}

func (mm *MockMachine) GetLabel() string {
    return mm.Label
}

func (mm *MockMachine) GetMode() string {
    return mm.Mode
}

func (mm *MockMachine) GetSudoPassword() string {
    return mm.SudoPassword
}

func (mm *MockMachine) GetDomain() string {
    return mm.Domain
}

func (mm *MockMachine) GetPort() string {
    return mm.Port
}

func (mm *MockMachine) GetUsername() string {
    return mm.Username
}
type MockTask struct {
    Summary string
    Task string
}

func (mt *MockTask) GetSummary() string {
    return mt.Summary
}

func (mt *MockTask) Serialize() string {
    return mt.Task
}


func TestExecute(t *testing.T) {
    c := &Command{}
    cases := []struct{
        name string
        e I_Executor
        err error
    }{
        {"no error", &MockExecutorSucc{}, nil},
        {"error", &MockExecutorFail{}, &exec.ExitError{}},
    }
    teardownCommandTestCase := setupCommandTestCase(t)
    defer teardownCommandTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownCommandSubTest := setupCommandSubTest(t)
            defer teardownCommandSubTest(t)
            err := c.Execute(tc.e)
            if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
                t.Fatalf("expected result %v, but got %v", tc.err, err)
            }
        })
    }
}

func TestGetPresentation(t *testing.T) {
    cases := []struct{
        name string
        t assignment.I_Task
        m assignment.I_Machine
        result []string
    }{
        {
            name: "SSHKEY mode",
            t: &MockTask{Task:"docker version && docker ps -a"},
            m: &MockMachine{Mode:"SSHKEY", Label:"mycomputer"},
            result: []string{"ssh", "-t", "mycomputer", "/bin/bash", "-c", "'docker version && docker ps -a'"},
        },
        {
            name: "USERPASS mode",
            t: &MockTask{Task:"docker version && docker ps -a"},
            m: &MockMachine{
                Mode:"USERPASS",
                Domain:"mycomputer",
                Port:"1112",
                Username:"me",
                SudoPassword: "mypass",
            },
            result: []string{"sshpass", "-p", "mypass", "ssh", "-t", "-p", "1112", "me@mycomputer", "/bin/bash", "-c", "'docker version && docker ps -a'"},
        },

        {
            name: "other mode",
            t: &MockTask{Task:"docker version && docker ps -a"},
            m: &MockMachine{Mode:"others", Label: "mycomputer"},
            result: []string{"/bin/bash", "-c", "'docker version && docker ps -a'"},
        },
    }
    teardownCommandTestCase := setupCommandTestCase(t)
    defer teardownCommandTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownCommandSubTest := setupCommandSubTest(t)
            defer teardownCommandSubTest(t)

            cmd := &Command{Machine: tc.m, Task: tc.t}
            presentation := cmd.GetPresentation()

            if !reflect.DeepEqual(presentation, tc.result) {
                t.Fatalf("expected result %v, but got %v", tc.result, presentation)
            }
        })
    }

}

func TestGetTask(t *testing.T) {
    cases := []struct{
        name string
        t assignment.I_Task
        m assignment.I_Machine
        result string
    }{
        {
            name: "without sudo",
            t: &MockTask{Task:"docker version && docker ps -a"},
            m: &MockMachine{SudoPassword: "pass"},
            result: "'docker version && docker ps -a'",
        },
        {
            name: "with sudo",
            t: &MockTask{Task:"docker version && sudo iptables -L -nv"},
            m: &MockMachine{SudoPassword: "pass"},
            result: "'docker version && echo pass | sudo -S iptables -L -nv'",
        },
    }
    teardownCommandTestCase := setupCommandTestCase(t)
    defer teardownCommandTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownCommandSubTest := setupCommandSubTest(t)
            defer teardownCommandSubTest(t)

            cmd := &Command{Task: tc.t, Machine: tc.m}
            task := cmd.GetTask()

            if task != tc.result {
                t.Fatalf("expected result %v, but got %v", tc.result, task)
            }
        })
    }


}


func TestGetTitle(t *testing.T) {
    cases := []struct{
        name string
        t assignment.I_Task
        result string
    }{
        {
            name: "with summary",
            t: &MockTask{Summary:"Docker information"},
            result: "Docker information",
        },
        {
            name: "without summary",
            t: &MockTask{},
            result: "",
        },
    }
    teardownCommandTestCase := setupCommandTestCase(t)
    defer teardownCommandTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownCommandSubTest := setupCommandSubTest(t)
            defer teardownCommandSubTest(t)

            cmd := &Command{Task: tc.t}
            title := cmd.GetTitle()

            if !strings.Contains(title, tc.result) {
                t.Fatalf("expected result %v in %v, but not", tc.result, title)
            }
        })
    }


}
