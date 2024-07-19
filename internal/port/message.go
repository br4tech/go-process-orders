package port

import amqp "github.com/rabbitmq/amqp091-go"

type IMessaging interface {
	Consume(queueName string, handler func(amqp.Delivery)) error
	Publish(queueName, message string) error
}
