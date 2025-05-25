
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