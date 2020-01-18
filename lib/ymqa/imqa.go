package ymqa

import (
	"github.com/streadway/amqp"
)

type (
	IMqA interface {
		Publish(queue string, msg_text interface{}) error
		Consume(queue string, callback AMQSubCallBack) (cnt *amqp.Connection, err error)
		ConsumeAck(queue string, callback AMQSubCallBack) (cnt *amqp.Connection, err error)
	}
)
