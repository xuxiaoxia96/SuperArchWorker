package core

import (
	"SuperArchWorker/middleware/taskcontrol"
	"SuperArchWorker/vars"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"
	"time"
)


func handleTask(d amqp.Delivery, taskInfoType reflect.Type){
	if taskInfoType == nil{
		logrus.Errorf("[DoTask][Get Type From FuncPool] No such Controller -- %s", d.RoutingKey)
		return
	}

	taskInfo := reflect.New(taskInfoType).Interface()
	err := json.Unmarshal(d.Body, taskInfo)
	if err != nil {
		logrus.Errorf("[DoTask][Deserial Msg] %s", err)
		return
	}

	taskId := reflect.ValueOf(taskInfo).Elem().FieldByName("TaskId").String()
	// set status running
	taskcontrol.SchedulerTaskControl.UpdateTaskControlStatus(taskId, taskcontrol.TaskRunning)

	handleMethod := reflect.ValueOf(taskInfo).MethodByName("Handle")
	res := handleMethod.Call( make([]reflect.Value , 0) )

	result := res[0].String()
	if result == ""{
		// set status exception
		taskcontrol.SchedulerTaskControl.UpdateTaskControlStatus(taskId, taskcontrol.TaskException)
	}else {
		taskcontrol.SchedulerTaskControl.SaveResultToDB(taskId, d.RoutingKey, result)
		// set status finish
		taskcontrol.SchedulerTaskControl.UpdateTaskControlStatus(taskId, taskcontrol.TaskFinish)
	}

}

func GoDoTask(input chan vars.InputForm, group *sync.WaitGroup){
	defer group.Done()
	for {
		select {
		case inputForm, ok := <-input:
			if !ok {
				return
			} else {
				start := time.Now()
				handleTask(inputForm.Msg, inputForm.TaskType)
				costTime := time.Since(start).Seconds()
				taskcontrol.SchedulerTaskControl.UpdateTaskTypeCostTime(inputForm.Msg.RoutingKey, costTime)
			}
		}
	}
}

