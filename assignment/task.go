package assignment



import (
    "strings"
)

/*
    Define interface: I_Task
*/
type I_Task interface {
    Serialize() string
    GetSummary() string
}

/*
    Define class: Task implementing I_Task
*/
type Task struct {
    Topic string
    Tasks []string
}


/*
    Generate task executable command
*/
func (t *Task) Serialize() string {
    return strings.Join(t.Tasks, " && ")
}


/*
    Get description of task
*/
func (t *Task) GetSummary() string {
    return t.Topic
}



