// controller/reservation_controller.go
package controller

// import (
// 	"fmt"
// 	"golang/Golang-Design/CarRental/dtos"
// 	"golang/Golang-Design/CarRental/service"
// )

// type ReservationController struct {
// 	service *service.ReservationService
// }

// func NewReservationController(service *service.ReservationService) *ReservationController {
// 	return &ReservationController{service: service}
// }

// func (rc *ReservationController) ReserveCar(req dtos.ReservationRequest) {
// 	resp, err := rc.service.MakeReservation(req)
// 	if err != nil {
// 		fmt.Println("Reservation failed:", err)
// 		return
// 	}
// 	fmt.Printf("Reservation confirmed!", resp)
// }

// func (rc *ReservationController) CancelReservation(id string) {
// 	rc.service.CancelReservation(id)
// 	fmt.Println("Reservation cancelled.")
// }
