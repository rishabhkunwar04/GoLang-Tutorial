package interfaces

type PaymentProcessor interface {
	Pay(amount float64) bool
}
