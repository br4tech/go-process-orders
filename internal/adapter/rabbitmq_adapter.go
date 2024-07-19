package adapter

import (
	"fmt"
	"log"

	"github.com/br4tech/go-process-orders/internal/port"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQAdapter(amqpURI string) port.IMessaging {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	return &RabbitMQAdapter{
		conn:    conn,
		channel: ch,
	}
}

// Publish publica uma mensagem na fila especificada
func (r *RabbitMQAdapter) Publish(queueName, message string) error {
	// Declara a fila dentro da função Publish
	q, err := r.channel.QueueDeclare(
		queueName, // Nome da fila
		false,     // Durável
		false,     // Exclusiva
		false,     // Auto-delete
		false,     // Exclusive
		nil,       // Argumentos extras
	)

	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	err = r.channel.Publish(
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	log.Printf(" [x] Published message to queue %s: %s\n", queueName, message)
	return nil
}

// Consume consome mensagens da fila especificada
func (r *RabbitMQAdapter) Consume(queueName string, handler func(amqp.Delivery)) error {
	msgs, err := r.channel.Consume(
		queueName, // Nome da fila
		"",        // Consumer tag
		true,      // Auto-ack
		false,     // Exclusive
		false,     // No-local
		false,     // No-wait
		nil,       // Args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received message from queue %s: %s\n", queueName, d.Body)
			handler(d)
		}
	}()

	log.Printf(" [*] Waiting for messages on queue %s. To exit press CTRL+C\n", queueName)
	<-make(chan bool) // Bloqueia a execução para manter o consumidor ativo

	return nil
}

// Close fecha a conexão e o canal com o RabbitMQ
func (r *RabbitMQAdapter) Close() error {
	if err := r.channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}
	if err := r.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}
