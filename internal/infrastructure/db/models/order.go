package models

type Order struct {
	ID            uint
	OrderStatus   string
	ClientID      int
	PaymentMethod string
	Deleted       bool     `gorm:"default:false"`
	Tickets       []Ticket `gorm:"many2many:order_tickets;"`
}
