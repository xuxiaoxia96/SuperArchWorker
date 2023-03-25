package taskcontrol

import (
	"SuperArchWorker/middleware/postgres"
	"SuperArchWorker/middleware/redis"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

const (
	TaskControlHkey = "TaskControl"
	TaskTypeHkey = "TaskType"
)

const (
	TaskWaiting = 0
	TaskRunning = 1
	TaskFinish = 2
	TaskException = 9
)

type TaskControl struct {
	Status 			int64
}

var SchedulerTaskControl TaskControl

type TaskType struct {
	AverCostTime 	float64
	MostCostTime	float64
}

func (taskControl *TaskControl) UpdateTaskTypeCostTime(module string, costTime float64){
	taskTypeStr := redis.GetValFromHsetBykey(TaskTypeHkey, module)
	taskType := TaskType{}
	json.Unmarshal([]byte(taskTypeStr), &taskType)
	if costTime > taskType.MostCostTime{
		taskType.MostCostTime = costTime
	}
	taskType.AverCostTime = (taskType.AverCostTime + costTime) / 2

	newTaskTypeStr,_ := json.Marshal(taskType)
	newTaskType := make(map[string]interface{})
	newTaskType[module] = string(newTaskTypeStr)

	redis.UpdateHset(TaskTypeHkey, newTaskType)
}

func (taskControl *TaskControl) UpdateTaskControlStatus(taskId string, status int64) {
	taskControl.Status = status
	newTaskControlStr, _ := json.Marshal(taskControl)
	newTaskControl := make(map[string]interface{})
	newTaskControl[taskId] = string(newTaskControlStr)
	redis.UpdateHset(TaskControlHkey, newTaskControl)
}

func (taskControl TaskControl) SaveResultToDB(requestId, module, result string){
	pdb := postgres.GetPostgresClient()
	if pdb == nil{
		return
	}

	sql := `INSERT INTO "TaskResult" (request_id, module, result) VALUES ($1, $2, $3)`
	_, err := pdb.Exec(sql, requestId, module, result)
	if err != nil {
		logrus.Errorf("[TaskControl Module][SaveResultToDB][Insert Data] %s", err)
		return 
	}
	logrus.Infof("[TaskControl Module][SaveResultToDB] Insert Successfully! requestId: %s, module: %s", requestId, module)
	pdb.Close()
}

type TaskModuleCommon struct {
	TaskId		string
}