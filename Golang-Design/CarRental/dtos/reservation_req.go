package dtos

import "time"

type ReservationRequest struct {
	CustomerName  string
	ContactInfo   string
	LicenseNumber string
	CarMake       string
	CarModel      string
	StartDate     time.Time
	EndDate       time.Time
}
