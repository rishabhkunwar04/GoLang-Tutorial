package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang/Golang-Design/CarRental/dtos"
	"golang/Golang-Design/CarRental/interfaces"
	"golang/Golang-Design/CarRental/repository"
)

type ReservationService struct {
	carRepo          *repository.CarRepo
	reservationRepo  *repository.ReservationRepository
	paymentProcessor interfaces.PaymentProcessor
}

func NewReservationService(carRepo *repository.CarRepo, resRepo *repository.ReservationRepository, pp interfaces.PaymentProcessor) *ReservationService {
	return &ReservationService{carRepo, resRepo, pp}
}

func (s *ReservationService) generateID() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return "RES" + hex.EncodeToString(bytes)
}

func (s *ReservationService) MakeReservation(req dtos.ReservationRequest) (*dtos.ReservationResponse, error) {
	cars := s.carRepo.FindAvilableCar(req.CarMake, req.CarModel)
	if len(cars) == 0 {
		return nil, fmt.Errorf("no available cars found")
	}

	selectedCar := cars[0]

	days := req.EndDate.Sub(req.StartDate).Hours() / 24
	price := selectedCar.PricePerDay * days
	fmt.Println(price)

	// if !s.paymentProcessor.Pay()(price) {
	// 	return nil, fmt.Errorf("payment failed")
	// }

	// reservation := &models.Reservation{
	// 	ReservationID: s.generateID(),
	// 	Customer: &models.Customer{
	// 		Name:                 req.CustomerName,
	// 		ContactInfo:          req.ContactInfo,
	// 		DriversLicenseNumber: req.LicenseNumber,
	// 	},
	// 	Car:        selectedCar,
	// 	StartDate:  req.StartDate,
	// 	EndDate:    req.EndDate,
	// 	TotalPrice: price,
	// }

	// s.reservationRepo.Add(reservation)
	// selectedCar.SetAvilability()(false)

	return &dtos.ReservationResponse{
		// ReservationID: reservation.RevervationID,
		// TotalPrice:    reservation.PriceAmount,
		Status: "Confirmed",
	}, nil
}

func (s *ReservationService) CancelReservation(reservationID string) {
	resList := s.reservationRepo.GetAll()
	for _, res := range resList {
		if res.RevervationID == reservationID {
			res.Car.SetAvilability(true)
			s.reservationRepo.Delete(reservationID)
			return
		}
	}
}
