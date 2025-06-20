package models

import "time"

type Reservation struct {
	RevervationID string
	Car           *Car
	Customer      *Customer
	PriceAmount   float64
	StartDate     time.Time
	EndDate       time.Time
}
