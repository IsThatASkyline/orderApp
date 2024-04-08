package models

type Concert struct {
	ID      uint
	Title   string   `json:"title"`
	Tickets []Ticket `json:"tickets"`
}
