package main

//******* Before  open closed principle *******//
/*
type Discount struct{}

func (d *Discount) CalculateDiscount(discountType string, Amount int) float64 {
	if discountType == "regular" {
		return 0.2 * float64(Amount)
	}
	if discountType == "premium" {
		return 0.5 * float64(Amount)
	}
	return float64(Amount)
}
func main() {
	calc := Discount{}
	discount := calc.CalculateDiscount("regular", 100)
	fmt.Println(discount)

}
*/

// ****AFTER OCP ****** //

import (
	"fmt"
)

// Discount interface
type Discount interface {
	Calculate(amount float64) float64
}

// RegularDiscount struct
type RegularDiscount struct{}

func (r RegularDiscount) Calculate(amount float64) float64 {
	return amount * 0.1
}

// PremiumDiscount struct
type PremiumDiscount struct{}

func (p PremiumDiscount) Calculate(amount float64) float64 {
	return amount * 0.2
}

// DiscountCalculator struct
type DiscountCalculator struct{}

func (dc DiscountCalculator) CalculateDiscount(d Discount, amount float64) float64 {
	return d.Calculate(amount)
}

// Main function
func main() {
	regular := RegularDiscount{}
	premium := PremiumDiscount{}

	calculator := DiscountCalculator{}

	regularDiscount := calculator.CalculateDiscount(regular, 100)
	premiumDiscount := calculator.CalculateDiscount(premium, 100)

	fmt.Println("Regular Discount:", regularDiscount)
	fmt.Println("Premium Discount:", premiumDiscount)
}
