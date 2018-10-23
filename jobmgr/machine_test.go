package jobmgr


import (
    "testing"
)


func setupTestCase(t *testing.T) func(t *testing.T) {
    t.Log("setup test case")
    return func(t *testing.T) {
        t.Log("teardown test case")
    }
}

func setupSubTest(t *testing.T) func(t *testing.T) {
    t.Log("setup sub test")
    return func(t *testing.T) {
        t.Log("teardown sub test")
    }
}

func TestMachine(t *testing.T) {
    cases := []struct {
        name string
        eLabel string
        eSudoPassword string
        eMode string
        eUsername string
        eDomain string
        ePort string
    }{
        {"empty", "", "", "", "", "",""},
        {"with empty mode and pass", "Asian", "", "", "","",""},
        {"with empty pass", "Asian","","USERPASS", "", "", ""},
        {"with empty mode", "Sun", "password", "", "", "", ""},
        {"with label, mode, pass", "Sun", "password", "SSHKEY", "", "", ""},
        {"with all", "Sun", "password", "USERPASS", "me", "127.0.0.1", "22"},
        {"with empty port", "Sun", "password", "USERPASS", "me", "127.0.0.1", ""},
    }

    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownSubTest := setupSubTest(t)
            defer teardownSubTest(t)

            m := Machine{Label:tc.eLabel, SudoPassword: tc.eSudoPassword, Mode: tc.eMode, Domain: tc.eDomain, Port:tc.ePort, Username: tc.eUsername}
            label := m.GetLabel()
            if label != tc.eLabel {
                t.Fatalf("expected label %v, but got %v", tc.eLabel, label)
            }
            mode := m.GetMode()
            if mode != tc.eMode {
                t.Fatalf("expected mode %v, but got %v", tc.eMode, mode)
            }
            pass := m.GetSudoPassword()
            if pass != tc.eSudoPassword {
                t.Fatalf("expected sudopassword %v, but got %v", tc.eSudoPassword, pass)
            }
            domain := m.GetDomain()
            if domain != tc.eDomain {
                t.Fatalf("expected domain %v, but got %v", tc.eDomain, domain)
            }
            username := m.GetUsername()
            if username != tc.eUsername {
                t.Fatalf("expected username %v, but got %v", tc.eUsername, username)
            }
            port := m.GetPort()
            if port != tc.ePort {
                t.Fatalf("expected port %v, but got %v", tc.ePort, port)
            }


        })
    }
}
