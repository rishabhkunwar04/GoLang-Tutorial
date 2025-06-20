package main

/*
type PaymentProcessor interface {
	Pay(amount float64)
}

type LegacyPayment struct{}

func (l *LegacyPayment) MakePayment(money float64) {
	fmt.Printf("Paid using legacy system: ₹%.2f\n", money)
}

type LegacyAdapter struct {
	legacy *LegacyPayment
}

func (a *LegacyAdapter) Pay(amount float64) {
	a.legacy.MakePayment(amount) // Adapts to new interface
}

func main() {
	var processor PaymentProcessor

	// Using the adapter to wrap legacy system
	processor = &LegacyAdapter{legacy: &LegacyPayment{}}

	// Now client uses the new interface
	processor.Pay(1000.00)
}

*/
