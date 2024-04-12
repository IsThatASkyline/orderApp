package dto

import (
	"time"
)

type Order struct {
	OrderStatus   string    `json:"orderStatus"`
	PaymentMethod string    `json:"paymentMethod"`
	ClientID      int       `json:"clientID"`
	OrderID       int       `json:"OrderID"`
	Tickets       []Ticket  `json:"tickets"`
	CreatedAt     time.Time `json:"createdAt"`
	TotalPrice    float64   `json:"totalPrice"`
}
