package main

import (
	"SuperArchWorker/conf"
	"SuperArchWorker/middleware/register"
)

func main() {
	conf.InitConfig()
	register.SchedulerRegister.Init()

	register.SchedulerRegister.ReceiveFromMQAndDo()
}
