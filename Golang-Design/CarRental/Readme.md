## Car Rental
1. The car rental system should allow customers to browse and reserve available cars for specific dates.
2. Each car should have details such as make, model, year, license plate number, and rental price per day.
3. Customers should be able to search for cars based on various criteria, such as car type, price range, and availability.
4. The system should handle reservations, including creating, modifying, and canceling reservations.
5. The system should keep track of the availability of cars and update their status accordingly.
6. The system should handle customer information, including name, contact details, and driver's license 7. information.
7. The system should handle payment processing for reservations.
8. The system should be able to handle concurrent reservations and ensure data consistency.



```go
// dto/reservation_dto.go
package dto

import "time"

type ReservationRequest struct {
	CustomerName     string
	ContactInfo      string
	LicenseNumber    string
	CarMake          string
	CarModel         string
	StartDate        time.Time
	EndDate          time.Time
}

type ReservationResponse struct {
	ReservationID string
	TotalPrice    float64
	Status        string
}

// models/car.go
package models

import "sync"

type Car struct {
	Make              string
	Model             string
	Year              int
	LicensePlate      string
	RentalPricePerDay float64
	Available         bool
	Mu                sync.Mutex
}

func (c *Car) IsAvailable() bool {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Available
}

func (c *Car) SetAvailable(available bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Available = available
}

// models/customer.go

type Customer struct {
	Name                 string
	ContactInfo          string
	DriversLicenseNumber string
}
// models/reservation.go
import "time"

type Reservation struct {
	ReservationID string
	Customer      *Customer
	Car           *Car
	StartDate     time.Time
	EndDate       time.Time
	TotalPrice    float64
}
// interfaces/payment_processor.go
package interfaces

type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}
type CreditCardPaymentProcessor struct{}

func (p *CreditCardPaymentProcessor) ProcessPayment(amount float64) bool {
	return true
}


// repository/car_repository.go
package repository

import (
	"strings"
	"sync"
	"time"

	"carrentalsystem/models"
)

type CarRepository struct {
	cars map[string]*models.Car
	mu   sync.RWMutex
}

func NewCarRepository() *CarRepository {
	return &CarRepository{
		cars: make(map[string]*models.Car),
	}
}

func (r *CarRepository) AddCar(car *models.Car) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars[car.LicensePlate] = car
}

func (r *CarRepository) GetAll() []*models.Car {
	r.mu.RLock()
	defer r.mu.RUnlock()
	cars := make([]*models.Car, 0)
	for _, c := range r.cars {
		cars = append(cars, c)
	}
	return cars
}

func (r *CarRepository) FindAvailableCars(make, model string, isAvailableFunc func(*models.Car) bool) []*models.Car {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var results []*models.Car
	for _, car := range r.cars {
		if strings.EqualFold(car.Make, make) &&
			strings.EqualFold(car.Model, model) &&
			isAvailableFunc(car) {
			results = append(results, car)
		}
	}
	return results
}

// repository/reservation_repository.go
package repository

import (
	"carrentalsystem/models"
	"sync"
)

type ReservationRepository struct {
	reservations map[string]*models.Reservation
	mu           sync.RWMutex
}

func NewReservationRepository() *ReservationRepository {
	return &ReservationRepository{
		reservations: make(map[string]*models.Reservation),
	}
}

func (r *ReservationRepository) Add(reservation *models.Reservation) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reservations[reservation.ReservationID] = reservation
}

func (r *ReservationRepository) GetAll() []*models.Reservation {
	r.mu.RLock()
	defer r.mu.RUnlock()
	all := make([]*models.Reservation, 0)
	for _, res := range r.reservations {
		all = append(all, res)
	}
	return all
}

func (r *ReservationRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.reservations, id)
}


// service/reservation_service.go
package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"carrentalsystem/dto"
	"carrentalsystem/interfaces"
	"carrentalsystem/models"
	"carrentalsystem/repository"
)

type ReservationService struct {
	carRepo         *repository.CarRepository
	reservationRepo *repository.ReservationRepository
	paymentProcessor interfaces.PaymentProcessor
}

func NewReservationService(carRepo *repository.CarRepository, resRepo *repository.ReservationRepository, pp interfaces.PaymentProcessor) *ReservationService {
	return &ReservationService{carRepo, resRepo, pp}
}

func (s *ReservationService) generateID() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return "RES" + hex.EncodeToString(bytes)
}

func (s *ReservationService) MakeReservation(req dto.ReservationRequest) (*dto.ReservationResponse, error) {
	cars := s.carRepo.FindAvailableCars(req.CarMake, req.CarModel, func(car *models.Car) bool {
		return car.IsAvailable()
	})

	if len(cars) == 0 {
		return nil, fmt.Errorf("no available cars found")
	}

	selectedCar := cars[0]

	days := req.EndDate.Sub(req.StartDate).Hours() / 24
	price := selectedCar.RentalPricePerDay * days

	if !s.paymentProcessor.ProcessPayment(price) {
		return nil, fmt.Errorf("payment failed")
	}

	reservation := &models.Reservation{
		ReservationID: s.generateID(),
		Customer: &models.Customer{
			Name:                 req.CustomerName,
			ContactInfo:          req.ContactInfo,
			DriversLicenseNumber: req.LicenseNumber,
		},
		Car:        selectedCar,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		TotalPrice: price,
	}

	s.reservationRepo.Add(reservation)
	selectedCar.SetAvailable(false)

	return &dto.ReservationResponse{
		ReservationID: reservation.ReservationID,
		TotalPrice:    reservation.TotalPrice,
		Status:        "Confirmed",
	}, nil
}

func (s *ReservationService) CancelReservation(reservationID string) {
	resList := s.reservationRepo.GetAll()
	for _, res := range resList {
		if res.ReservationID == reservationID {
			res.Car.SetAvailable(true)
			s.reservationRepo.Delete(reservationID)
			return
		}
	}
}

// controller/reservation_controller.go
package controller

import (
	"carrentalsystem/dto"
	"carrentalsystem/service"
	"fmt"
)

type ReservationController struct {
	service *service.ReservationService
}

func NewReservationController(service *service.ReservationService) *ReservationController {
	return &ReservationController{service: service}
}

func (rc *ReservationController) ReserveCar(req dto.ReservationRequest) {
	resp, err := rc.service.MakeReservation(req)
	if err != nil {
		fmt.Println("Reservation failed:", err)
		return
	}
	fmt.Printf("Reservation confirmed! ID: %s, Price: %.2f\n", resp.ReservationID, resp.TotalPrice)
}

func (rc *ReservationController) CancelReservation(id string) {
	rc.service.CancelReservation(id)
	fmt.Println("Reservation cancelled.")
}


package main

import (
	"carrentalsystem/controller"
	"carrentalsystem/dto"
	"carrentalsystem/interfaces"
	"carrentalsystem/models"
	"carrentalsystem/repository"
	"carrentalsystem/service"
	"time"
)

func main() {
	// Setup dependencies
	carRepo := repository.NewCarRepository()
	reservationRepo := repository.NewReservationRepository()
	paymentProcessor := &interfaces.CreditCardPaymentProcessor{}
	reservationService := service.NewReservationService(carRepo, reservationRepo, paymentProcessor)
	reservationController := controller.NewReservationController(reservationService)

	// Seed some cars
	car1 := &models.Car{
		Make:              "Toyota",
		Model:             "Camry",
		Year:              2020,
		LicensePlate:      "ABC123",
		RentalPricePerDay: 1500,
		Available:         true,
	}
	car2 := &models.Car{
		Make:              "Honda",
		Model:             "City",
		Year:              2021,
		LicensePlate:      "XYZ789",
		RentalPricePerDay: 1300,
		Available:         true,
	}
	carRepo.AddCar(car1)
	carRepo.AddCar(car2)

	// Sample reservation request
	start := time.Now().AddDate(0, 0, 1)
	end := start.AddDate(0, 0, 3)

	req := dto.ReservationRequest{
		CustomerName:  "Rishabh Kunwar",
		ContactInfo:   "rishabh@example.com",
		LicenseNumber: "DL-0420201123456",
		CarMake:       "Toyota",
		CarModel:      "Camry",
		StartDate:     start,
		EndDate:       end,
	}

	// Make reservation
	reservationController.ReserveCar(req)

	// Cancel reservation example (if needed)
	// reservationController.CancelReservation("RESxxxxxx")
}

```