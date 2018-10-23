package jobmgr



import (
    "strings"
)

/*
    Define interface: I_Task
*/
type I_Task interface {
    Serialize() string
    GetTopic() string
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
func (t *Task) GetTopic() string {
    return t.Topic
}



