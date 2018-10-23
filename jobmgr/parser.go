package jobmgr


import (
    "encoding/json"
)

/*
    Parse the machine raw data, returning I_Machine
*/
func ParseMachineData(raw []byte) I_Machine {
    var m Machine
    json.Unmarshal(raw, &m)
    return &m
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
