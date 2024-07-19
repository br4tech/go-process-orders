package entities

import "time"

type Order struct {
	Id        string
	ProductId string
	Quantity  int
	CreatedAt time.Time
}
