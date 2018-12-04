package jobmgr
import (
  "encoding/json"
  "log"
  "github.com/gorilla/websocket"
)


/*
    Define interface: I_Mgr for handling ssh job and websocket communication
*/
type I_Mgr interface {
    SetConn(*websocket.Conn)
    ExecuteJob()
}

/*
    Implement I_Mgr
*/
type Mgr struct {
    Conn *websocket.Conn
    Job I_Job
}

/*
    Set websocket connection
*/
func (mgr *Mgr) SetConn(conn *websocket.Conn) {
    mgr.Conn = conn
}

/*
    Create job from job description, which is defined as jobmgr.Job
*/
func (mgr *Mgr) CreateJob(jobData []byte) error {
    tmpJob := struct {
        Machine Machine
        Tasks []Task
    }{Machine: Machine{}, Tasks: []Task{}}
    err := json.Unmarshal(jobData, &tmpJob)

    job := &Job{Machine: &tmpJob.Machine, Tasks: []I_Task{}}
    for index, _ := range tmpJob.Tasks {
        job.Tasks = append(job.Tasks, &tmpJob.Tasks[index])
    }
    mgr.Job = job
    return err
}

/*
    Get job description from websocket
*/
func (mgr *Mgr) getConf() []byte {
    _, jobData, err := mgr.Conn.ReadMessage()
    if err != nil {
        log.Fatal("GetConf", err)
    }
    return jobData
}

/*
    Execute ssh job
*/
func (mgr *Mgr) ExecuteJob() {
    mgr.CreateJob(mgr.getConf())
    mgr.Job.Execute(mgr.Conn)
}
