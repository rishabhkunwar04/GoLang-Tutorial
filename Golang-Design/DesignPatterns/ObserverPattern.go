package main

/*
1. Define the Observer Interface

package main

import "fmt"

type Observer interface {
	Update(status string)
}
2. Concrete Observers

type EmailService struct{}

func (e EmailService) Update(status string) {
	fmt.Println("📧 EmailService: Booking status changed to", status)
}

type SMSService struct{}

func (s SMSService) Update(status string) {
	fmt.Println("📱 SMSService: Booking status changed to", status)
}
3. Subject Interface and Implementation

type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify()
	SetStatus(status string)
}

type Booking struct {
	observers []Observer
	status    string
}

func (b *Booking) Register(o Observer) {
	b.observers = append(b.observers, o)
}

func (b *Booking) Remove(o Observer) {
	for i, obs := range b.observers {
		if obs == o {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *Booking) Notify() {
	for _, o := range b.observers {
		o.Update(b.status)
	}
}

func (b *Booking) SetStatus(status string) {
	b.status = status
	b.Notify()
}
4. Test / Main

func main() {
	booking := &Booking{}

	email := EmailService{}
	sms := SMSService{}

	booking.Register(email)
	booking.Register(sms)

	booking.SetStatus("Confirmed")

	// Output:
	// 📧 EmailService: Booking status changed to Confirmed
	// 📱 SMSService: Booking status changed to Confirmed
}
*/
