

## payment system
```go
Q. Design a payment system like amazon have when we click on pay it give option to select payment method and direct us to 3rd party system. Design database and table schema for it and also design api endpoints we will be needing. Implement above feature using interfaces,services and dependency injection ?package main

import (
	"fmt"
)

// Database Schema (in comments):
/*
Tables:
1. Users
   - id (PK)
   - name
   - email

2. Orders
   - id (PK)
   - user_id (FK)
   - total_amount
   - status (PENDING, PAID)

3. Payments
   - id (PK)
   - order_id (FK)
   - payment_method (CARD, PAYPAL, UPI)
   - status (INITIATED, SUCCESS, FAILED)
   - transaction_reference
*/

// PaymentMethod defines supported payment types
type PaymentMethod string

const (
	Card   PaymentMethod = "CARD"
	Paypal PaymentMethod = "PAYPAL"
	UPI    PaymentMethod = "UPI"
)

// PaymentRequest holds payment info from user
type PaymentRequest struct {
	OrderID       int
	Amount        float64
	Method        PaymentMethod
	UserID        int
	PaymentDetail map[string]string // cardNo, upiId etc.
}

// PaymentService interface
// Each implementation will represent 3rd party gateway like Stripe, PayPal, etc.
type PaymentService interface {
	Pay(req PaymentRequest) (string, error) // returns transaction reference
}

// CardPaymentService implements PaymentService
type CardPaymentService struct{}

func (c *CardPaymentService) Pay(req PaymentRequest) (string, error) {
	fmt.Println("Processing card payment...")
	// Simulate third party call
	return "txn_card_123456", nil
}

// PaypalPaymentService implements PaymentService
type PaypalPaymentService struct{}

func (p *PaypalPaymentService) Pay(req PaymentRequest) (string, error) {
	fmt.Println("Processing PayPal payment...")
	return "txn_paypal_123456", nil
}

// UPIPaymentService implements PaymentService
type UPIPaymentService struct{}

func (u *UPIPaymentService) Pay(req PaymentRequest) (string, error) {
	fmt.Println("Processing UPI payment...")
	return "txn_upi_123456", nil
}

// PaymentProcessor uses dependency injection
// Selects proper service based on request

type PaymentProcessor struct {
	cardService   PaymentService
	paypalService PaymentService
	upiService    PaymentService
}

func NewPaymentProcessor(card PaymentService, paypal PaymentService, upi PaymentService) *PaymentProcessor {
	return &PaymentProcessor{
		cardService:   card,
		paypalService: paypal,
		upiService:    upi,
	}
}

func (p *PaymentProcessor) ProcessPayment(req PaymentRequest) error {
	var service PaymentService
	switch req.Method {
	case Card:
		service = p.cardService
	case Paypal:
		service = p.paypalService
	case UPI:
		service = p.upiService
	default:
		return fmt.Errorf("unsupported payment method")
	}

	ref, err := service.Pay(req)
	if err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	fmt.Printf("Payment successful, transaction ref: %s\n", ref)
	// Here, store payment record to DB
	return nil
}

func main() {
	processor := NewPaymentProcessor(&CardPaymentService{}, &PaypalPaymentService{}, &UPIPaymentService{})

	req := PaymentRequest{
		OrderID: 1,
		Amount: 100.0,
		Method: Card,
		UserID: 101,
		PaymentDetail: map[string]string{
			"cardNumber": "1234-5678-9012-3456",
		},
	}

	err := processor.ProcessPayment(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Sample API Endpoints (in REST):
/*
POST /api/payments/initiate
Request Body:
{
  "orderId": 1,
  "userId": 101,
  "amount": 100.0,
  "method": "CARD",
  "details": {
    "cardNumber": "1234-5678-9012-3456"
  }
}
*/

```
