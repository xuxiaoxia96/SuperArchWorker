package register

import (
	"SuperArchWorker/conf"
	"SuperArchWorker/core"
	"SuperArchWorker/core/ai/painter"
	"SuperArchWorker/vars"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

type Register struct {
	ModulePool map[string]reflect.Type
}

var SchedulerRegister Register

func (register *Register)Init()  {
	register.ModulePool = make(map[string]reflect.Type)
	register.ModulePool["ai.painter.txt2image"] = reflect.TypeOf(painter.Txt2Image{})
	register.ModulePool["ai.painter.image2image"] = reflect.TypeOf(painter.Image2Image{})
	register.ModulePool["ai.painter.imagescore"] = reflect.TypeOf(painter.ImageScore{})
	register.ModulePool["ai.painter.image2txt"] = reflect.TypeOf(painter.Image2Txt{})
	register.ModulePool["ai.painter.imagevqa"] = reflect.TypeOf(painter.ImageVQA{})
}

func (register *Register)ReceiveFromMQAndDo(){
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		conf.Cfg.RabbitMQ.Username, conf.Cfg.RabbitMQ.Password, conf.Cfg.RabbitMQ.Host, conf.Cfg.RabbitMQ.Port))

	if err != nil{
		logrus.Errorf("[DoTask][RabbitMQ Dial] %s", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		logrus.Errorf("[DoTask][Open Channel] %s", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		conf.Cfg.RabbitMQ.Exchange, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil{
		logrus.Errorf("[DoTask][Declare Exchange] %s", err)
		return
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil{
		logrus.Errorf("[DoTask][Declare Queue] %s", err)
		return
	}

	for _, s := range reflect.ValueOf(SchedulerRegister.ModulePool).MapKeys() {
		routeKey := s.Interface().(string)
		logrus.Infof("Binding queue %s to exchange with routing key %s",
			q.Name, routeKey)
		err = ch.QueueBind(
			q.Name,       // queue name
			routeKey,            // routing key
			conf.Cfg.RabbitMQ.Exchange, // exchange
			false,
			nil)
		if err != nil{
			logrus.Errorf("[DoTask][Bind Key] key: %s, error: %s", s, err)
			return
		}
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil{
		logrus.Errorf("[DoTask][Register Consumer] %s", err)
		return
	}

	fetchWg := sync.WaitGroup{}
	inputChan := make(chan vars.InputForm, *vars.ProcessNum*4)
	for i := 0; i < *vars.ProcessNum; i++ {
		go core.GoDoTask(inputChan, &fetchWg) // start all workers
		fetchWg.Add(1)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			inputForm := vars.InputForm{
				d,
				SchedulerRegister.ModulePool[d.RoutingKey],
			}
			inputChan <- inputForm
		}
	}()

	logrus.Infof(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	close(inputChan)
	fetchWg.Wait()
}
