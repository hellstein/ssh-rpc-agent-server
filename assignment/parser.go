package assignment


import (
    "io/ioutil"
    "encoding/json"
    "fmt"
)

/*
    Parse the machine configuration file, returning []I_Machine
*/
func ParseMachineFile(filename string) []I_Machine {
    r := []I_Machine{}
    raw, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err.Error())
        return r
    }

    var ms []Machine
    json.Unmarshal(raw, &ms)
    for i, _ := range ms {
        r = append(r, &ms[i])
    }
    return r
}

/*
    Parse the task configuration file, returning []I_Task
*/
func ParseTaskFile(filename string) []I_Task {
    r := []I_Task{}
    raw, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err.Error())
        return r
    }

    var ts []Task
    json.Unmarshal(raw, &ts)
    for i, _ := range ts {
        r = append(r, &ts[i])
    }
    return r
}
