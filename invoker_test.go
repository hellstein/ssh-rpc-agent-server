package main


import (
    "testing"
    "reflect"
)


func TestInitProcessor(t *testing.T) {
    cases := []struct {
        name string
        tfile string
        mfile string
    }{
        {"empty", "", ""},
        {"without mfile", "t.json",""},
        {"without tfile", "","m.json"},
        {"with all", "t.json", "m.json"},
    }

    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            teardownSubTest := setupSubTest(t)
            defer teardownSubTest(t)

            p := InitProcessor(tc.tfile, tc.mfile)
            ptype := reflect.TypeOf(p)
            pkind := ptype.Kind()
            pname := ptype.String()
            if pkind != reflect.Ptr {
                t.Fatalf("expected type %v, but got %v", reflect.Ptr, pkind)
            }
            if pname != "*main.Processor" {
                t.Fatalf("expected object %v, but got %v", "*main.Processor", pname)
            }

        })
    }

}


// TODO
func TestInvoker(t *testing.T) {
}
