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

### singleton design pattern
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

## Builder design patterm
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

## Google Doc Design
```go
package main

import (
	"fmt"
)

// --- User Struct ---
type User struct {
	Username string
	UserID   string
}

// --- Permission Enum ---
type Permission int

const (
	READ Permission = iota
	WRITE
	OWNER
)

func (p Permission) String() string {
	return [...]string{"READ", "WRITE", "OWNER"}[p]
}

// --- Document Struct ---
type Document struct {
	Content string
	DocName string
	DocMap  map[string]Permission
}

func NewDocument(user User, content string, docName string) *Document {
	return &Document{
		Content: content,
		DocName: docName,
		DocMap:  map[string]Permission{user.UserID: OWNER},
	}
}

func (d *Document) GrantAccess(user User, perm Permission) {
	d.DocMap[user.UserID] = perm
}

func (d *Document) RevokeAccess(user User) {
	delete(d.DocMap, user.UserID)
}

func (d *Document) WriteContent(user User, content string) {
	if perm, ok := d.DocMap[user.UserID]; ok && (perm == OWNER || perm == WRITE) {
		d.Content += content
		fmt.Println("Content written by", user.Username)
	} else {
		fmt.Println("Permission denied to WRITE!!")
	}
}

func (d *Document) ReadContent(user User) {
	if perm, ok := d.DocMap[user.UserID]; ok && (perm == OWNER || perm == READ) {
		fmt.Println("Content read by", user.Username)
		fmt.Println("-- Content --\n" + d.Content)
	} else {
		fmt.Println("Permission denied to READ!!")
	}
}

func (d *Document) DeleteContent(user User) {
	if perm, ok := d.DocMap[user.UserID]; ok && perm == OWNER {
		d.Content = ""
		fmt.Println("Content deleted by", user.Username)
	} else {
		fmt.Println("Permission denied to DELETE!!")
	}
}

// --- Main ---
func main() {
	user1 := User{"Rishabh", "rishabh1"}
	user2 := User{"Sumit", "sumit1"}
	user3 := User{"Sushant", "sushant1"}

	doc := NewDocument(user1, "initial content !! ", "myDoc")

	fmt.Println(doc.Content)
	doc.WriteContent(user1, "content added by owner")
	fmt.Println(doc.Content)

	doc.GrantAccess(user2, READ)
	doc.RevokeAccess(user2)

	doc.ReadContent(user2) // Should say permission denied
}
```

## Simple Parking Lot
```go
1. Vehicle
Number: string – Vehicle registration number

Type: VehicleType – Enum (Bike, Car, Truck)

2. ParkingSlot
SlotNumber: int – Unique identifier for the slot

IsOccupied: bool – Flag to indicate if the slot is occupied

Vehicle: *Vehicle – Pointer to the parked vehicle

SlotType: VehicleType – Type of vehicle this slot supports

3. ParkingLot
Slots: []*ParkingSlot – List of all slots in the parking lot

4. Enum: VehicleType
Constants:

Bike

Car

Truck

📘 UML Class Diagram (Textual Representation)
plaintext
Copy
Edit
+------------------+
|     Vehicle      |
+------------------+
| - Number: string |
| - Type: VehicleType |
+------------------+

+-----------------------+
|    ParkingSlot        |
+-----------------------+
| - SlotNumber: int     |
| - IsOccupied: bool    |
| - Vehicle: *Vehicle   |
| - SlotType: VehicleType |
+-----------------------+

+-----------------------+
|     ParkingLot        |
+-----------------------+
| - Slots: []*ParkingSlot |
+-----------------------+
| + ParkVehicle(v *Vehicle): (int, error) |
| + RemoveVehicle(slotNumber int): error |
+-----------------------+

+-------------------+
|   VehicleType     |
+-------------------+
| + Bike            |
| + Car             |
| + Truck           |
+-------------------+

package main

import (
	"errors"
	"fmt"
)

// VehicleType defines types of vehicles
type VehicleType int

const (
	Bike VehicleType = iota
	Car
	Truck
)

func (v VehicleType) String() string {
	switch v {
	case Bike:
		return "Bike"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	default:
		return "Unknown"
	}
}

type Vehicle struct {
	Number string
	Type   VehicleType
}

type ParkingSlot struct {
	SlotNumber int
	IsOccupied bool
	Vehicle    *Vehicle
	SlotType   VehicleType
}

type ParkingLot struct {
	Slots []*ParkingSlot
}

func NewParkingLot(numSlots int) *ParkingLot {
	slots := make([]*ParkingSlot, numSlots)
	for i := 0; i < numSlots; i++ {
		var slotType VehicleType
		switch {
		case i%3 == 0:
			slotType = Bike
		case i%3 == 1:
			slotType = Car
		default:
			slotType = Truck
		}
		slots[i] = &ParkingSlot{
			SlotNumber: i + 1,
			SlotType:   slotType,
		}
	}
	return &ParkingLot{Slots: slots}
}

func (pl *ParkingLot) ParkVehicle(vehicle *Vehicle) (int, error) {
	for _, slot := range pl.Slots {
		if !slot.IsOccupied && slot.SlotType == vehicle.Type {
			slot.IsOccupied = true
			slot.Vehicle = vehicle
			fmt.Printf("Vehicle %s parked at slot %d\n", vehicle.Number, slot.SlotNumber)
			return slot.SlotNumber, nil
		}
	}
	return -1, errors.New("No available slot for vehicle type")
}

func (pl *ParkingLot) RemoveVehicle(slotNumber int) error {
	if slotNumber < 1 || slotNumber > len(pl.Slots) {
		return errors.New("Invalid slot number")
	}
	slot := pl.Slots[slotNumber-1]
	if !slot.IsOccupied {
		return errors.New("Slot is already empty")
	}
	fmt.Printf("Vehicle %s removed from slot %d\n", slot.Vehicle.Number, slot.SlotNumber)
	slot.Vehicle = nil
	slot.IsOccupied = false
	return nil
}

func main() {
	lot := NewParkingLot(10)
	v1 := &Vehicle{Number: "KA-01-HH-1234", Type: Car}
	v2 := &Vehicle{Number: "KA-01-HH-9999", Type: Bike}
	v3 := &Vehicle{Number: "KA-01-BB-0001", Type: Truck}

	lot.ParkVehicle(v1)
	lot.ParkVehicle(v2)
	lot.ParkVehicle(v3)

	lot.RemoveVehicle(2)
}

```

## Parking lot having ticket system

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

// VehicleType defines types of vehicles
type VehicleType int

const (
	Bike VehicleType = iota
	Car
	Truck
)

func (v VehicleType) String() string {
	switch v {
	case Bike:
		return "Bike"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	default:
		return "Unknown"
	}
}

type Vehicle struct {
	Number string
	Type   VehicleType
}

type ParkingSlot struct {
	SlotNumber int
	IsOccupied bool
	Vehicle    *Vehicle
	SlotType   VehicleType
	StartTime  time.Time
}

type Ticket struct {
	SlotNumber int
	Vehicle    *Vehicle
	EntryTime  time.Time
	ExitTime   time.Time
	Fee        float64
}

type ParkingLot struct {
	Slots []*ParkingSlot
}

func NewParkingLot(numSlots int) *ParkingLot {
	slots := make([]*ParkingSlot, numSlots)
	for i := 0; i < numSlots; i++ {
		var slotType VehicleType
		switch {
		case i%3 == 0:
			slotType = Bike
		case i%3 == 1:
			slotType = Car
		default:
			slotType = Truck
		}
		slots[i] = &ParkingSlot{
			SlotNumber: i + 1,
			SlotType:   slotType,
		}
	}
	return &ParkingLot{Slots: slots}
}

func (pl *ParkingLot) ParkVehicle(vehicle *Vehicle) (*Ticket, error) {
	for _, slot := range pl.Slots {
		if !slot.IsOccupied && slot.SlotType == vehicle.Type {
			slot.IsOccupied = true
			slot.Vehicle = vehicle
			slot.StartTime = time.Now()
			fmt.Printf("Vehicle %s parked at slot %d\n", vehicle.Number, slot.SlotNumber)
			return &Ticket{
				SlotNumber: slot.SlotNumber,
				Vehicle:    vehicle,
				EntryTime:  slot.StartTime,
			}, nil
		}
	}
	return nil, errors.New("No available slot for vehicle type")
}

func (pl *ParkingLot) RemoveVehicle(slotNumber int) (*Ticket, error) {
	if slotNumber < 1 || slotNumber > len(pl.Slots) {
		return nil, errors.New("Invalid slot number")
	}
	slot := pl.Slots[slotNumber-1]
	if !slot.IsOccupied {
		return nil, errors.New("Slot is already empty")
	}
	exitTime := time.Now()
	duration := exitTime.Sub(slot.StartTime)
	fee := calculateFee(duration, slot.SlotType)

	ticket := &Ticket{
		SlotNumber: slot.SlotNumber,
		Vehicle:    slot.Vehicle,
		EntryTime:  slot.StartTime,
		ExitTime:   exitTime,
		Fee:        fee,
	}

	fmt.Printf("Vehicle %s removed from slot %d\n", slot.Vehicle.Number, slot.SlotNumber)
	fmt.Printf("Duration: %.2f hours, Fee: %.2f\n", duration.Hours(), fee)

	slot.Vehicle = nil
	slot.IsOccupied = false
	slot.StartTime = time.Time{}

	return ticket, nil
}

func calculateFee(duration time.Duration, vType VehicleType) float64 {
	hours := duration.Hours()
	rate := 0.0
	switch vType {
	case Bike:
		rate = 10
	case Car:
		rate = 20
	case Truck:
		rate = 30
	}
	return rate * hours
}

func main() {
	lot := NewParkingLot(10)
	v1 := &Vehicle{Number: "KA-01-HH-1234", Type: Car}
	v2 := &Vehicle{Number: "KA-01-HH-9999", Type: Bike}
	v3 := &Vehicle{Number: "KA-01-BB-0001", Type: Truck}

	t1, _ := lot.ParkVehicle(v1)
	t2, _ := lot.ParkVehicle(v2)
	t3, _ := lot.ParkVehicle(v3)

	time.Sleep(2 * time.Second) 

	lot.RemoveVehicle(t1.SlotNumber)
}

```
## Social Media
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type NotificationType string

const (
	FriendRequest         NotificationType = "FriendRequest"
	FriendRequestAccepted NotificationType = "FriendRequestAccepted"
	Like                  NotificationType = "Like"
	Comment               NotificationType = "Comment"
	Mention               NotificationType = "Mention"
)

type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	ProfilePic  string
	Bio         string
	Interests   []string
	Friends     map[int]bool
	Posts       []int
	Notifications []Notification
}

type Post struct {
	ID        int
	UserID    int
	Content   string
	Images    []string
	Videos    []string
	Timestamp time.Time
	Likes     map[int]bool
	Comments  []Comment
}

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	Content   string
	Timestamp time.Time
}

type Notification struct {
	ID        int
	UserID    int
	Type      NotificationType
	Content   string
	Timestamp time.Time
}

type SocialNetworkingService struct {
	Users   map[int]*User
	Posts   map[int]*Post
	Mutex   sync.Mutex
	NextUID int
	NextPID int
	NextCID int
	NextNID int
}

var instance *SocialNetworkingService
var once sync.Once

func GetSNSInstance() *SocialNetworkingService {
	once.Do(func() {
		instance = &SocialNetworkingService{
			Users: make(map[int]*User),
			Posts: make(map[int]*Post),
			Mutex: sync.Mutex{},
		}
	})
	return instance
}

func (sns *SocialNetworkingService) RegisterUser(name, email, password string) *User {
	sns.Mutex.Lock()
	defer sns.Mutex.Unlock()
	user := &User{
		ID:       sns.NextUID,
		Name:     name,
		Email:    email,
		Password: password,
		Friends:  make(map[int]bool),
	}
	sns.Users[user.ID] = user
	sns.NextUID++
	return user
}

func (sns *SocialNetworkingService) CreatePost(userID int, content string, images, videos []string) *Post {
	sns.Mutex.Lock()
	defer sns.Mutex.Unlock()
	post := &Post{
		ID:        sns.NextPID,
		UserID:    userID,
		Content:   content,
		Images:    images,
		Videos:    videos,
		Timestamp: time.Now(),
		Likes:     make(map[int]bool),
	}
	sns.Posts[post.ID] = post
	sns.Users[userID].Posts = append(sns.Users[userID].Posts, post.ID)
	sns.NextPID++
	return post
}

func (sns *SocialNetworkingService) SendFriendRequest(senderID, receiverID int) {
	notification := Notification{
		ID:        sns.NextNID,
		UserID:    receiverID,
		Type:      FriendRequest,
		Content:   fmt.Sprintf("%s sent you a friend request.", sns.Users[senderID].Name),
		Timestamp: time.Now(),
	}
	sns.Users[receiverID].Notifications = append(sns.Users[receiverID].Notifications, notification)
	sns.NextNID++
}

func (sns *SocialNetworkingService) AcceptFriendRequest(userID, friendID int) {
	sns.Users[userID].Friends[friendID] = true
	sns.Users[friendID].Friends[userID] = true
	notification := Notification{
		ID:        sns.NextNID,
		UserID:    friendID,
		Type:      FriendRequestAccepted,
		Content:   fmt.Sprintf("%s accepted your friend request.", sns.Users[userID].Name),
		Timestamp: time.Now(),
	}
	sns.Users[friendID].Notifications = append(sns.Users[friendID].Notifications, notification)
	sns.NextNID++
}

func (sns *SocialNetworkingService) LikePost(userID, postID int) {
	post := sns.Posts[postID]
	post.Likes[userID] = true
	owner := sns.Users[post.UserID]
	notification := Notification{
		ID:        sns.NextNID,
		UserID:    owner.ID,
		Type:      Like,
		Content:   fmt.Sprintf("%s liked your post.", sns.Users[userID].Name),
		Timestamp: time.Now(),
	}
	owner.Notifications = append(owner.Notifications, notification)
	sns.NextNID++
}

func (sns *SocialNetworkingService) CommentOnPost(userID, postID int, content string) {
	comment := Comment{
		ID:        sns.NextCID,
		UserID:    userID,
		PostID:    postID,
		Content:   content,
		Timestamp: time.Now(),
	}
	sns.Posts[postID].Comments = append(sns.Posts[postID].Comments, comment)
	sns.NextCID++
}

func (sns *SocialNetworkingService) GetNewsFeed(userID int) []Post {
	var feed []Post
	for friendID := range sns.Users[userID].Friends {
		for _, postID := range sns.Users[friendID].Posts {
			feed = append(feed, *sns.Posts[postID])
		}
	}
	for _, postID := range sns.Users[userID].Posts {
		feed = append(feed, *sns.Posts[postID])
	}
	return feed // Simplified, no sort for now
}

func main() {
	sns := GetSNSInstance()
	u1 := sns.RegisterUser("Alice", "alice@example.com", "pass123")
	u2 := sns.RegisterUser("Bob", "bob@example.com", "pass456")

	sns.SendFriendRequest(u1.ID, u2.ID)
	sns.AcceptFriendRequest(u2.ID, u1.ID)

	p1 := sns.CreatePost(u1.ID, "Hello from Alice!", nil, nil)
	sns.LikePost(u2.ID, p1.ID)
	sns.CommentOnPost(u2.ID, p1.ID, "Nice post!")

	feed := sns.GetNewsFeed(u1.ID)
	for _, post := range feed {
		fmt.Println("Post:", post.Content)
	}
}

```
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

## Movie booking System
```go
+-------------+
|    User     |
+-------------+
| ID          |
| Name        |
| Email       |
+-------------+

+-------------+
|    Movie    |
+-------------+
| ID          |
| Title       |
| Duration    |
| Genre       |
+-------------+

+---------------+
|   Theater     |
+---------------+
| ID            |
| Name          |
| Location      |
| Screens[]     |
+---------------+

+-------------+
|   Screen     |
+-------------+
| ID           |
| Name         |
| TotalSeats   |
| Seats[]      |
+-------------+

+-------------+
|   Seat       |
+-------------+
| ID           |
| Row, Column  |
| IsBooked     |
+-------------+

+-------------+
|    Show      |
+-------------+
| ID           |
| Movie        |
| Screen       |
| StartTime    |
| EndTime      |
+-------------+

+----------------+
|    Booking      |
+----------------+
| ID              |
| User            |
| Show            |
| BookedSeats[]   |
| BookingTime     |
| Amount          |
+----------------+

package main

import (
	"errors"
	"fmt"
	"time"
)

type SeatStatus int

const (
	Available SeatStatus = iota
	Booked
)

type Seat struct {
	Row     int
	Number  int
	Status  SeatStatus
	UserID  string
}

type Show struct {
	MovieName   string
	Theater     string
	Screen      int
	StartTime   time.Time
	Seats       [][]*Seat
}

type Booking struct {
	BookingID string
	UserID    string
	ShowID    string
	Seats     []string
	BookedAt  time.Time
}

type BookingSystem struct {
	Shows    map[string]*Show
	Bookings map[string]*Booking
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		Shows:    make(map[string]*Show),
		Bookings: make(map[string]*Booking),
	}
}

func (bs *BookingSystem) CreateShow(id, movie, theater string, screen int, rows, cols int, start time.Time) {
	seats := make([][]*Seat, rows)
	for i := range seats {
		seats[i] = make([]*Seat, cols)
		for j := 0; j < cols; j++ {
			seats[i][j] = &Seat{Row: i, Number: j, Status: Available}
		}
	}
	bs.Shows[id] = &Show{
		MovieName: movie,
		Theater:   theater,
		Screen:    screen,
		StartTime: start,
		Seats:     seats,
	}
}

func (bs *BookingSystem) BookSeats(userID, showID string, seatRequests [][2]int) (*Booking, error) {
	show, ok := bs.Shows[showID]
	if !ok {
		return nil, errors.New("show not found")
	}
	seatIDs := []string{}
	for _, req := range seatRequests {
		r, c := req[0], req[1]
		if show.Seats[r][c].Status != Available {
			return nil, fmt.Errorf("seat (%d,%d) already booked", r, c)
		}
	}
	for _, req := range seatRequests {
		r, c := req[0], req[1]
		show.Seats[r][c].Status = Booked
		show.Seats[r][c].UserID = userID
		seatIDs = append(seatIDs, fmt.Sprintf("%d-%d", r, c))
	}
	bookingID := fmt.Sprintf("BKG-%d", time.Now().UnixNano())
	booking := &Booking{
		BookingID: bookingID,
		UserID:    userID,
		ShowID:    showID,
		Seats:     seatIDs,
		BookedAt:  time.Now(),
	}
	bs.Bookings[bookingID] = booking
	return booking, nil
}

func main() {
	bs := NewBookingSystem()
	bs.CreateShow("SHOW123", "Inception", "PVR", 1, 5, 5, time.Now().Add(2*time.Hour))

	booking, err := bs.BookSeats("user1", "SHOW123", [][2]int{{0, 0}, {0, 1}})
	if err != nil {
		fmt.Println("Booking failed:", err)
	} else {
		fmt.Println("Booking successful:", booking)
	}
}

```