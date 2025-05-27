package main

/*
import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
}

type UpiPayment struct{}

func (u UpiPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using UPI\n", amount)
}

type Booking struct {
	Amount          float64
	PaymentStrategy PaymentStrategy
}

func (b *Booking) ProcessPayment() {
	b.PaymentStrategy.Pay(b.Amount)
}

func main() {
	booking := &Booking{
		Amount:          299.99,
		PaymentStrategy: UpiPayment{},
	}
	booking.ProcessPayment()

	// Change strategy at runtime
	booking.PaymentStrategy = CreditCardPayment{}
	booking.ProcessPayment()
}

*/
