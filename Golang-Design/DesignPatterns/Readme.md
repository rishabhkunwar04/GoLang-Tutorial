## Singleton design pattern
- The Singleton Design Pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance.
**usecases**:
1. *Logging Service:* A single logger instance shared across modules
2. *Thread Pool Manager:*	One shared pool to avoid resource overhead
3. *Database Connection Pool*:	Manage all DB connections centrally
```Golang

1. Lazy Initialization (Not Thread-Safe)

package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Singleton struct{}

var instance *Singleton
var once sync.Once

// getInstance is similar to Java's getInstance()
func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := getInstance()
	s2 := getInstance()

	fmt.Println("Are s1 and s2 the same?", s1 == s2) // true
}

 2. Eager Initialization

package main

import "fmt"

// Singleton struct - unexported to prevent direct instantiation
type Singleton struct{}

// Eagerly created singleton instance
var instance = &Singleton{}

func GetInstance() *Singleton {
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println("Are s1 and s2 the same?", s1 == s2) // Output: true
}


 3. Thread safe lazy initialization
 package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Singleton struct{}

var instance *Singleton
var once sync.Once

// getInstance is similar to Java's getInstance()
func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := getInstance()
	s2 := getInstance()

	fmt.Println("Are s1 and s2 the same?", s1 == s2) // true
}


 4. Double-Checked Locking

package singleton

import (
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	lock     sync.Mutex
)

func GetInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}
 4. Enum Singleton Equivalent in Go

Go doesn’t have enums with behavior like Java, but you can achieve the same effect using a constant struct instance.

package singleton

type enumSingleton struct{}

var EnumInstance = &enumSingleton{}

func (e *enumSingleton) DoSomething() {
	// logic here
}
```

 ## Builder Design Pattern 
* it is a creational design pattern that solves problems related to creating complex objects with multiple configurations. It provides a systematic way to construct an object step-by-step while ensuring that the construction process is independent of the object's representation.

**when to use builder design pattern**
1. Object has many optional parameters
2. Object construction is complex or involves multiple steps

**disadvantage** : extra boiler plate code added for obj creation

**Immutability and Builder**
Builder pattern promotes immutability because:
- All fields are set inside the builder, not changed after construction.
- The built object can have only getter methods (no setters).
- Ensures thread-safety and predictable behavior in concurrent environments.

```go
package main

import "fmt"

// Car is the product being built
type Car struct {
	Make    string
	Model   string
	Year    int
	Color   string
	Sunroof bool
}

func (c Car) String() string {
	return fmt.Sprintf("Car{Make: %q, Model: %q, Year: %d, Color: %q, Sunroof: %t}",
		c.Make, c.Model, c.Year, c.Color, c.Sunroof)
}

// CarBuilder builds a Car
type CarBuilder struct {
	make    string
	model   string
	year    int
	color   string
	sunroof bool
}

// NewCarBuilder is the constructor for CarBuilder
func NewCarBuilder(make, model string) *CarBuilder {
	return &CarBuilder{
		make:  make,
		model: model,
	}
}

func (b *CarBuilder) Year(year int) *CarBuilder {
	b.year = year
	return b
}

func (b *CarBuilder) Color(color string) *CarBuilder {
	b.color = color
	return b
}

func (b *CarBuilder) Sunroof(sunroof bool) *CarBuilder {
	b.sunroof = sunroof
	return b
}

// Build constructs the final Car
func (b *CarBuilder) Build() Car {
	return Car{
		Make:    b.make,
		Model:   b.model,
		Year:    b.year,
		Color:   b.color,
		Sunroof: b.sunroof,
	}
}

// Main usage
func main() {
	car := NewCarBuilder("Toyota", "Camry").
		Year(2022).
		Color("Blue").
		Sunroof(true).
		Build()

	fmt.Println(car)
}

```

## factory design pattern

It is a creational design pattern used to create objects without exposing the instantiation logic to the client, and instead using a common interface. It's a key pattern when you want to encapsulate object creation, particularly when object construction is complex or varies based on some conditions

* Used when you need a single factory method to create related objects based on certain inputs or conditions.
* We use factory design pattern when there is super class and multiple subclass, we want to use subclasses
based on input or requirement
*  It  Provides loose coupling and more robust code
* it encapsulate object creation logic
- **Example:** In a movie booking system, creating different types of seats (Normal, Premium) with dynamic pricing and status.
- **Disadvantage**
    1. Complexity of code
    2. Hidden object creation logic so harder to debug


```go
package main

import "fmt"

// Shape interface
type Shape interface {
	Draw()
}

// Square struct
type Square struct{}

func (s Square) Draw() {
	fmt.Println("Drawing Square")
}

// Circle struct
type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Drawing Circle")
}

// ShapeFactory struct
type ShapeFactory struct{}

func (sf ShapeFactory) GetShape(input string) Shape {
	switch input {
	case "Circle":
		return Circle{}
	case "Square":
		return Square{}
	default:
		return Square{} // default case
	}
}

// Test (main function)
func main() {
	factory := ShapeFactory{}
	shape := factory.GetShape("Square")
	shape.Draw()
}

```

## Strategy Design Pattern
- The Strategy Design Pattern is a behavioral design pattern that enables selecting an algorithm's behavior at runtime
**when to use**:
1. Multiple algorithms for a task
2. Behavior changes at runtime

```go
1. Define the Strategy Interface

type PaymentStrategy interface {
	Pay(amount float64)
}
2. Implement Concrete Strategies

type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
}

type UpiPayment struct{}

func (u UpiPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using UPI\n", amount)
}
3. Booking Context Using Strategy

type Booking struct {
	Amount          float64
	PaymentStrategy PaymentStrategy
}

func (b *Booking) ProcessPayment() {
	b.PaymentStrategy.Pay(b.Amount)
}
4. Usage

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
```

## Observer Design Pattern

* The Observer design pattern is a behavioral design pattern, used to create a one-to-many dependency between objects so that when one object (the subject) changes its state, all its dependents (observers) are notified and updated automatically. 

**when to use**: Common in event-driven systems, UI frameworks, messaging systems, and real-time notifications.
##### Components of the Observer Design Pattern
* Subject: The subject maintains a list of observers and notifies them of state changes.
* Observer: The observer interface defines the contract for concrete observer classes.
* ConcreteSubject: A class that implements the subject interface and manages the observers.
* ConcreteObserver: A class that implements the observer interface and receives notifications.

```go
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
```
