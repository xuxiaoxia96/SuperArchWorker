package vars

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"reflect"
)

const (
	VersionInfo = "0.0.1"
)

type InputForm struct {
	Msg			amqp.Delivery
	TaskType 	reflect.Type
}
