package ymq

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type (
	IMq interface {
		Publish(nodeID int64, topic string, msg interface{}) error
		PublishQos(nodeID int64, topic string, qos byte, msg interface{}) error
		Subscribe(nodeID int64, topic string, callback MQTT.MessageHandler) (cnt *Mqtt, err error)
	}
)
