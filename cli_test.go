package main

import (
    "testing"
    "reflect"
    "github.com/urfave/cli"
//    "fmt"
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

func TestInitFlag(t *testing.T) {
    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    teardownSubTest := setupSubTest(t)
    defer teardownSubTest(t)


    flags := InitFlags()
    if len(flags) != 2 {
        t.Fatalf("expected number of flags %v, but got %v", 2, len(flags))
    }
}


func TestInitAppAction(t *testing.T) {
    teardownTestCase := setupTestCase(t)
    defer teardownTestCase(t)

    teardownSubTest := setupSubTest(t)
    defer teardownSubTest(t)
    
    argsHandler := func (c *cli.Context) error {
        return nil
    }
    invoker := &Invoker{}
    
    f := InitAppAction(argsHandler, invoker)
    ftype := reflect.TypeOf(f).Kind()
    if ftype != reflect.Func {
        t.Fatalf("expected type of result is %v, but got %v", reflect.Func, ftype)
    }

}


// TODO
/*
func TestArgsHandler(t *testing.T) {
    c := &cli.Context{}
    c.Set("machinefile", "")
    fmt.Println(c)
}
*/



// TODO
func TestInitCli(t *testing.T) {
}
