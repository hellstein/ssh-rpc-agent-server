package jobmgr


import (
    "testing"
)

func TestMachineFileParser(t *testing.T) {
    cases := []struct {
        name string
        filename string
        eLen int
    }{
        {"without filename", "", 0},
        {"file doesn't exist", "xxx.json", 0},
        {"file exist but incorrect", "incorrect.json", 0},
        {"correct file", "correct.machine.json", 4},
    }


    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownSubTest := setupSubTest(t)
            defer teardownSubTest(t)

            lms := len(ParseMachineFile(tc.filename))
            if lms != tc.eLen {
                t.Fatalf("expected length of machines is %v, but got %v", tc.eLen, lms)
            }
        })
    }

}


func TestTaskFileParser(t *testing.T) {
    cases := []struct {
        name string
        filename string
        eLen int
    }{
        {"without filename", "", 0},
        {"file doesn't exist", "xxx.json", 0},
        {"file exist but incorrect", "incorrect.json", 0},
        {"correct file", "correct.task.json", 5},
    }


    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownSubTest := setupSubTest(t)
            defer teardownSubTest(t)

            lts := len(ParseTaskFile(tc.filename))
            if lts != tc.eLen {
                t.Fatalf("expected length of tasks is %v, but got %v", tc.eLen, lts)
            }
        })
    }

}
