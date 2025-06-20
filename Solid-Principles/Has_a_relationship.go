package main

// import "fmt"

// type Engine struct {
// }

// func (e *Engine) Start() {
// 	fmt.Println("Engine started")

// }

// type Car struct {
// 	engine Engine
// }

// func (c *Car) startcar() {
// 	c.engine.Start()
// }

// func main() {
// 	c := Car{}
// 	c.startcar()

// }

// composition in golang
/*
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


*/
