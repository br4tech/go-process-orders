package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/br4tech/go-process-orders/internal/domain/entities"
	"github.com/br4tech/go-process-orders/internal/port"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderService struct {
	message port.IMessaging
}

func NeworderService(message port.IMessaging) *OrderService {
	return &OrderService{
		message: message,
	}
}

func (service *OrderService) CreateOrders() {
	for i := 1; i <= 10; i++ {
		order := entities.Order{
			Id:        fmt.Sprintf("order-%d", i),
			ProductId: fmt.Sprintf("product-%d", i),
			Quantity:  i,
			CreatedAt: time.Now(),
		}
		orderJson, _ := json.Marshal(order)
		if err := service.message.Publish("orders", string(orderJson)); err != nil {
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second)
	}
}

func (service *OrderService) ListAllOrders() ([]entities.Order, error) {
	var orders []entities.Order

	err := service.message.Consume("orders", func(d amqp.Delivery) {
		var order entities.Order
		if err := json.Unmarshal(d.Body, &order); err != nil {
			log.Printf("Error unmarshaling order: %s", err)
			return
		}
		orders = append(orders, order)
	})

	if err != nil {
		return nil, fmt.Errorf("failed to consume orders: %w", err)
	}

	return orders, nil
}
