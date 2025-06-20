package main

// import (
// 	"golang/Golang-Design/CarRental/controller"
// 	"golang/Golang-Design/CarRental/dtos"
// 	"golang/Golang-Design/CarRental/interfaces"
// 	"golang/Golang-Design/CarRental/models"
// 	"golang/Golang-Design/CarRental/repository"
// 	"golang/Golang-Design/CarRental/service"
// 	"time"
// )

// func main() {
// 	// Setup dependencies
// 	carRepo := repository.NewCarRepo()
// 	reservationRepo := repository.NewReservationRepository()
// 	paymentProcessor := &interfaces.CreditCardPaymentProcessor{}
// 	reservationService := service.NewReservationService(carRepo, reservationRepo, paymentProcessor)
// 	reservationController := controller.NewReservationController(reservationService)

// 	// Seed some cars

// 	car1 := &models.Car{
// 		Name:                 "Toyota",
// 		Model:                "Camry",
// 		Licence_Plate_Number: "ABC123",
// 		PricePerDay:          1500,
// 		Avilable:             true,
// 	}

// 	carRepo.AddCar(car1)

// 	// Sample reservation request
// 	start := time.Now().AddDate(0, 0, 1)
// 	end := start.AddDate(0, 0, 3)

// 	req := dtos.ReservationRequest{
// 		CustomerName:  "Rishabh Kunwar",
// 		ContactInfo:   "rishabh@example.com",
// 		LicenseNumber: "DL-0420201123456",
// 		CarMake:       "Toyota",
// 		CarModel:      "Camry",
// 		StartDate:     start,
// 		EndDate:       end,
// 	}

// 	// Make reservation
// 	reservationController.ReserveCar(req)

// 	// Cancel reservation example (if needed)
// 	// reservationController.CancelReservation("RESxxxxxx")
// }
