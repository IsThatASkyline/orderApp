package models

type Ticket struct {
	ID        uint
	Price     int
	Place     string
	IsSold    bool `gorm:"default:false"`
	ConcertID uint
}
