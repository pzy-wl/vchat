package yamq

import (
	"github.com/streadway/amqp"
)

func getCntOfRabbitMQ(url string) (*amqp.Connection, error) {
	//conn, err := amqp.Dial("amqp://root:password@localhost:5672/")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func publishWrap(ch *amqp.Channel, data *AMqData) error {

	return nil
}
func consumeWrap(queue string, f func(body *amqp.Delivery) error) error {
	return nil
}
