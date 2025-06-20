package interfaces

type CreditCardPaymentProcessor struct {
}

func (c *CreditCardPaymentProcessor) Pay(amount float64) bool {

	// payment logic if success return true else return false

	return true
}
