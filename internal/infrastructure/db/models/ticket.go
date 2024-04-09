package models

type Ticket struct {
	ID         uint
	EventTitle string
	Price      int
	Place      string
	IsSold     bool `gorm:"default:false"`
}
