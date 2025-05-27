package main

/*
import (
	"fmt"
)

// Component interface
type Coffee interface {
	GetCost() int
	GetDescription() string
}

// Concrete component
type SimpleCoffee struct{}

func (s *SimpleCoffee) GetCost() int {
	return 5
}

func (s *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

// Base decorator (embeds the component)
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) GetCost() int {
	return d.coffee.GetCost()
}

func (d *CoffeeDecorator) GetDescription() string {
	return d.coffee.GetDescription()
}

// Concrete decorators
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(c Coffee) Coffee {
	return &MilkDecorator{CoffeeDecorator{coffee: c}}
}

func (m *MilkDecorator) GetCost() int {
	return m.coffee.GetCost() + 2
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", Milk"
}

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(c Coffee) Coffee {
	return &SugarDecorator{CoffeeDecorator{coffee: c}}
}

func (s *SugarDecorator) GetCost() int {
	return s.coffee.GetCost() + 1
}

func (s *SugarDecorator) GetDescription() string {
	return s.coffee.GetDescription() + ", Sugar"
}

// --- Main client code ---
func main() {
	var coffee Coffee = &SimpleCoffee{}
	fmt.Println(coffee.GetDescription(), "=> $", coffee.GetCost())

	coffee = NewMilkDecorator(coffee)
	coffee = NewSugarDecorator(coffee)
	coffee = NewSugarDecorator(coffee) // Add extra sugar

	fmt.Println(coffee.GetDescription(), "=> $", coffee.GetCost())
}

*/
