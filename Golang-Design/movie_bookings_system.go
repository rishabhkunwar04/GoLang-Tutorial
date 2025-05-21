package main

import (
	"errors"
	"fmt"
	"time"
)

type SeatStatus int

const (
	Available SeatStatus = iota
	Booked
)

type Seat struct {
	Row    int
	Number int
	Status SeatStatus
	UserID string
}

type Show struct {
	MovieName string
	Theater   string
	Screen    int
	StartTime time.Time
	Seats     [][]*Seat
}

type Booking struct {
	BookingID string
	UserID    string
	ShowID    string
	Seats     []string
	BookedAt  time.Time
}

type BookingSystem struct {
	Shows    map[string]*Show
	Bookings map[string]*Booking
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		Shows:    make(map[string]*Show),
		Bookings: make(map[string]*Booking),
	}
}

func (bs *BookingSystem) CreateShow(id, movie, theater string, screen int, rows, cols int, start time.Time) {
	seats := make([][]*Seat, rows)
	for i := range seats {
		seats[i] = make([]*Seat, cols)
		for j := 0; j < cols; j++ {
			seats[i][j] = &Seat{Row: i, Number: j, Status: Available}
		}
	}
	bs.Shows[id] = &Show{
		MovieName: movie,
		Theater:   theater,
		Screen:    screen,
		StartTime: start,
		Seats:     seats,
	}
}

func (bs *BookingSystem) BookSeats(userID, showID string, seatRequests [][2]int) (*Booking, error) {
	show, ok := bs.Shows[showID]
	if !ok {
		return nil, errors.New("show not found")
	}
	seatIDs := []string{}
	for _, req := range seatRequests {
		r, c := req[0], req[1]
		if show.Seats[r][c].Status != Available {
			return nil, fmt.Errorf("seat (%d,%d) already booked", r, c)
		}
	}
	for _, req := range seatRequests {
		r, c := req[0], req[1]
		show.Seats[r][c].Status = Booked
		show.Seats[r][c].UserID = userID
		seatIDs = append(seatIDs, fmt.Sprintf("%d-%d", r, c))
	}
	bookingID := fmt.Sprintf("BKG-%d", time.Now().UnixNano())
	booking := &Booking{
		BookingID: bookingID,
		UserID:    userID,
		ShowID:    showID,
		Seats:     seatIDs,
		BookedAt:  time.Now(),
	}
	bs.Bookings[bookingID] = booking
	return booking, nil
}

func main() {
	bs := NewBookingSystem()
	bs.CreateShow("SHOW123", "Inception", "PVR", 1, 5, 5, time.Now().Add(2*time.Hour))

	booking, err := bs.BookSeats("user1", "SHOW123", [][2]int{{0, 0}, {0, 1}})
	if err != nil {
		fmt.Println("Booking failed:", err)
	} else {
		fmt.Println("Booking successful:", booking)
	}
}
