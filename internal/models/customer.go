package models

import "time"

// Customer is the customer model
type Customer struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	Password   string
	AddressLat float32
	AddressLon float32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
