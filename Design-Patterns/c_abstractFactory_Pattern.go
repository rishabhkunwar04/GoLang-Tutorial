package main

/*
import (
	"fmt"
	"strings"
)

///////////////////////
// Abstract Products //
///////////////////////

type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

////////////////////////
// Concrete Products  //
////////////////////////

type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

type Square struct{}

func (s Square) Draw() {
	fmt.Println("Drawing a Square")
}

type Red struct{}

func (r Red) Fill() {
	fmt.Println("Filling with Red")
}

type Blue struct{}

func (b Blue) Fill() {
	fmt.Println("Filling with Blue")
}

//////////////////////
// Abstract Factory //
//////////////////////

type AbstractFactory interface {
	GetShape() Shape
	GetColor() Color
}

//////////////////////
// Shape Factory    //
//////////////////////

type ShapeFactory struct {
	shapeType string
}

func (sf ShapeFactory) GetShape() Shape {
	switch strings.ToUpper(sf.shapeType) {
	case "CIRCLE":
		return Circle{}
	case "SQUARE":
		return Square{}
	default:
		return nil
	}
}

func (sf ShapeFactory) GetColor() Color {
	return nil // ShapeFactory doesn't create colors
}

//////////////////////
// Color Factory    //
//////////////////////

type ColorFactory struct {
	colorType string
}

func (cf ColorFactory) GetShape() Shape {
	return nil // ColorFactory doesn't create shapes
}

func (cf ColorFactory) GetColor() Color {
	switch strings.ToUpper(cf.colorType) {
	case "RED":
		return Red{}
	case "BLUE":
		return Blue{}
	default:
		return nil
	}
}

//////////////////////
// Factory Producer //
//////////////////////

func GetFactory(choice, typeName string) AbstractFactory {
	switch strings.ToUpper(choice) {
	case "SHAPE":
		return ShapeFactory{shapeType: typeName}
	case "COLOR":
		return ColorFactory{colorType: typeName}
	default:
		return nil
	}
}

//////////////////////
// Client Code      //
//////////////////////

func main() {
	// Get Shape Factory
	shapeFactory := GetFactory("SHAPE", "CIRCLE")
	if shape := shapeFactory.GetShape(); shape != nil {
		shape.Draw() // Output: Drawing a Circle
	}

	// Get Color Factory
	colorFactory := GetFactory("COLOR", "RED")
	if color := colorFactory.GetColor(); color != nil {
		color.Fill() // Output: Filling with Red
	}
}

*/
