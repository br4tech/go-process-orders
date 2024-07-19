package main

import (
	"fmt"
	"log"
	"time"

	"github.com/br4tech/go-process-orders/internal/adapter"
	"github.com/br4tech/go-process-orders/internal/domain/services"
)

func main() {
	amqpURI := "amqp://guest:guest@localhost:5672/"
	message := adapter.NewRabbitMQAdapter(amqpURI)
	service := services.NeworderService(message)

	service.CreateOrders()

	go func() {
		time.Sleep(5 * time.Second)

		orders, err := service.ListAllOrders()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Process orders:", orders)
	}()

	select {}
}
