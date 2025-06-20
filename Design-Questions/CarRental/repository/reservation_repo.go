// repository/reservation_repository.go
package repository

// import (
// 	"golang/Golang-Design/CarRental/models"
// 	"sync"
// )

// type ReservationRepository struct {
// 	reservations map[string]*models.Reservation
// 	mu           sync.RWMutex
// }

// func NewReservationRepository() *ReservationRepository {
// 	return &ReservationRepository{
// 		reservations: make(map[string]*models.Reservation),
// 	}
// }

// func (r *ReservationRepository) Add(reservation *models.Reservation) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	r.reservations[reservation.RevervationID] = reservation
// }

// func (r *ReservationRepository) GetAll() []*models.Reservation {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()
// 	all := make([]*models.Reservation, 0)
// 	for _, res := range r.reservations {
// 		all = append(all, res)
// 	}
// 	return all
// }

// func (r *ReservationRepository) Delete(id string) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	delete(r.reservations, id)
// }
