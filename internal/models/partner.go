package models

import "time"

// Partner is the partner model
type Partner struct {
	ID              int
	FirstName       string
	LastName        string
	Email           string
	AddressLat      int
	AddressLon      int
	OperatingRadius int
	Rating          int
	AccessLevel     int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
