package assignment


import (
    "testing"
    "strings"
)


func TestTask(t *testing.T) {
    cases := []struct {
        name string
        eTopic string
        eTasks []string
    }{
        {"null", "", nil},
        {"empty", "", []string{},},
        {"one", "list files", []string{"ls"},},
        {"more", "list files", []string{"ls", "ls -al", "ls -h",}},
    }

    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownSubTest := setupSubTest(t)
            defer teardownSubTest(t)

            m := Task{Topic:tc.eTopic, Tasks: tc.eTasks}
            summary := m.GetSummary()
            if summary != tc.eTopic {
                t.Fatalf("expected summary %v, but got %v", tc.eTopic, summary)
            }
            task := m.Serialize()
            eTask := strings.Join(tc.eTasks, " && ")
            if task != eTask {
                t.Fatalf("expected serialization %v, but got %v", eTask, task)
            }

        })
    }
}
