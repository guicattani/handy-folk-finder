package models

import "time"

// Partner is the partner model
type Partner struct {
	ID              int
	FirstName       string
	LastName        string
	Email           string
	AddressLat      float32
	AddressLon      float32
	OperatingRadius int
	Rating          int
	Experience      []string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
