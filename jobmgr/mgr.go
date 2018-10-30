package jobmgr
import (
  "encoding/json"
  "log"
  "github.com/gorilla/websocket"
)


type I_Mgr interface {
  CreateJob(conn *websocket.Conn) (I_Job, error)
  ExecuteJob(job I_Job, conn *websocket.Conn)
}


type Mgr struct {
}

func (mgr *Mgr) CreateJob(conn *websocket.Conn) (I_Job, error) {
  tmpJob := struct {
    Machine Machine
    Tasks []Task
  }{Machine: Machine{}, Tasks: []Task{}}
  _, jobData, err := conn.ReadMessage()
  if err != nil {
    log.Println("GetConf", err)
  }
  err = json.Unmarshal(jobData, &tmpJob)

  job := &Job{Machine: &tmpJob.Machine, Tasks: []I_Task{}}
  for index, _ := range tmpJob.Tasks {
    job.Tasks = append(job.Tasks, &tmpJob.Tasks[index])
  }
  return job, err
}

func (mgr *Mgr) ExecuteJob(job I_Job, conn *websocket.Conn) {
  job.Execute(conn)
}
