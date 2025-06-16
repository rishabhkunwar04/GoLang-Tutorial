package interfaces

type PaymentProcessor interface {
	Pay(amount int) bool
}
