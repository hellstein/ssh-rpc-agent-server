package jobmgr


import (
    "encoding/json"
)

/*
    Parse the machine raw data, returning []I_Machine
*/
func ParseMachineData(raw []byte) []I_Machine {
    r := []I_Machine{}
    var ms []Machine
    json.Unmarshal(raw, &ms)
    for i, _ := range ms {
        r = append(r, &ms[i])
    }
    return r
}

/*
    Parse the task raw data, returning []I_Task
*/
func ParseTaskData(raw []byte) []I_Task {
    r := []I_Task{}
    var ts []Task
    json.Unmarshal(raw, &ts)
    for i, _ := range ts {
        r = append(r, &ts[i])
    }
    return r
}
