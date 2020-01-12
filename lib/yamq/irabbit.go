package yamq

import (
	"github.com/streadway/amqp"
)

type (
	IMq interface {
		Publish(queue string, msg_text interface{}) error
		PublishAck(queue string, msg_text interface{}) error
		Consume(queue, callback func(body amqp.Delivery) error) (cnt *amqp.Connection, err error)
		ConsumeAck(queue, callback func(body amqp.Delivery) error) (cnt *amqp.Connection, err error)
	}
)
