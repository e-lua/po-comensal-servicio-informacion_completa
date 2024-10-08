package models

import (
	"github.com/streadway/amqp"
)

type MqttDB struct {
	amqp.Config
}

var MqttCN = GetConnMqtt()

func GetConnMqtt() *amqp.Connection {

	conn, error_connec_mqtt := amqp.Dial("amqp://edwardlopez:servermqtt@165.227.69.88:8888/")

	if error_connec_mqtt != nil {

		return nil
	}

	return conn
}
