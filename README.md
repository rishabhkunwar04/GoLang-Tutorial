### Has_a relation ship
```golang
package main

import "fmt"

type Engine struct {
}

func (e *Engine) Start() {
	fmt.Println("Engine started")

}

type Car struct {
	engine Engine
}

func (c *Car) startcar() {
	c.engine.Start()
}

func main() {
	c := Car{}
	c.startcar()

}

```

### composition in golang
```golang
package main

import "fmt"

// Define an interface
type Engine interface {
	Start()
}

// Concrete implementation of Engine
type PetrolEngine struct{}

func (p PetrolEngine) Start() {
	fmt.Println("Petrol engine started")
}

// Another implementation of Engine
type DieselEngine struct{}

func (d DieselEngine) Start() {
	fmt.Println("Diesel engine started")
}

// Car "has-a" Engine → composition
type Car struct {
	engine Engine // Interface, not a specific engine
}

func (c Car) StartCar() {
	fmt.Print("Starting the car: ")
	c.engine.Start()
}

func main() {
	petrolEngine := PetrolEngine{}
	dieselEngine := DieselEngine{}

	car1 := Car{engine: petrolEngine}
	car2 := Car{engine: dieselEngine}

	car1.StartCar() // Output: Starting the car: Petrol engine started
	car2.StartCar() // Output: Starting the car: Diesel engine started
}

```

### Open closed principle
```Golang
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

```
### Interface Segeration
```Golang
package main

import "fmt"

////  ****** Before interface segeration ********** //
/*
import "fmt"

type Printer interface {
	PrintDocument()
	ScanDocument()
	FaxDocument()
}

type SimplePrinter struct{}

func (sp SimplePrinter) PrintDocument() {
	fmt.Println("Printing document...")
}

func (sp SimplePrinter) ScanDocument() {
	// Not implemented
}

func (sp SimplePrinter) FaxDocument() {
	// Not implemented
}

func main() {
	var printer Printer = SimplePrinter{}
	printer.PrintDocument()
}

*/
//  ****** After interface segeration ********** //

// Segregated interfaces
type Printer interface {
	PrintDocument()
}

type Scanner interface {
	ScanDocument()
}

type Faxer interface {
	FaxDocument()
}

// Only implements what it needs
type SimplePrinter struct{}

func (sp SimplePrinter) PrintDocument() {
	fmt.Println("Printing document...")
}

// Another printer that does all
type AdvancedPrinter struct{}

func (ap AdvancedPrinter) PrintDocument() {
	fmt.Println("Advanced: Printing document...")
}

func (ap AdvancedPrinter) ScanDocument() {
	fmt.Println("Advanced: Scanning document...")
}

func (ap AdvancedPrinter) FaxDocument() {
	fmt.Println("Advanced: Faxing document...")
}

func main() {
	var p Printer = SimplePrinter{}
	p.PrintDocument()

	var ap Printer = AdvancedPrinter{}
	ap.PrintDocument()
}
```

### Dependency Inversion
```Golang
❌ DIP Violation Example in Go
go
Copy
Edit
package main

import "fmt"

// Low-Level Module
type EmailService struct{}

func (e EmailService) SendEmail(message string) {
	fmt.Println("Email sent:", message)
}

// High-Level Module
type Notification struct {
	emailService EmailService // Direct dependency on concrete class
}

func NewNotification() *Notification {
	return &Notification{
		emailService: EmailService{}, // Instantiating concrete type
	}
}

func (n *Notification) Notify(message string) {
	n.emailService.SendEmail(message)
}

// Main function
func main() {
	notification := NewNotification()
	notification.Notify("Hello, DIP!")
}


✅ DIP-Compliant Example in Go
go
Copy
Edit
package main

import "fmt"

// Abstraction
type NotificationService interface {
	SendMessage(message string)
}

// Low-Level Module 1: Email Service
type EmailService struct{}

func (e EmailService) SendMessage(message string) {
	fmt.Println("Email sent:", message)
}

// Low-Level Module 2: SMS Service
type SMSService struct{}

func (s SMSService) SendMessage(message string) {
	fmt.Println("SMS sent:", message)
}

// High-Level Module
type Notification struct {
	service NotificationService
}

// Constructor Injection
func NewNotification(service NotificationService) *Notification {
	return &Notification{service: service}
}

func (n *Notification) Notify(message string) {
	n.service.SendMessage(message)
}

// Main Function
func main() {
	emailService := EmailService{}
	smsService := SMSService{}

	emailNotification := NewNotification(emailService)
	emailNotification.Notify("Hello via Email!")

	smsNotification := NewNotification(smsService)
	smsNotification.Notify("Hello via SMS!")
}
```

### Dependency Injection
```golang
package main

import (
	"fmt"
)

type printer interface {
	print(message string)
}

type printMachine struct{}

func (p *printMachine) print(message string) {
	fmt.Println(message)
}

type app struct {
	printApp printer
}

// now dpendency injection via constructor
func NewApp(printApp2 printer) *app {
	return &app{printApp: printApp2}
}

func (a *app) run() {
	a.printApp.print("printing...")
}
func main() {
	printer := printMachine{}
	app := NewApp(&printer)
	app.run()

}
```