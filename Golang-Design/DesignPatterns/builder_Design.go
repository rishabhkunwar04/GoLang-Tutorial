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
