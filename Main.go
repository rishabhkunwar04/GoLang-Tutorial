package main

import "fmt"

type Shape interface {
	draw() string
}

type Circle struct{}

func (c *Circle) draw() string {
	return "circle class"
}

type Square struct{}

func (s *Square) draw() string {
	return "square class"
}

type ShapeFactory struct{}

func (sh *ShapeFactory) getShape(str string) Shape {
	if str == "circle" {
		return &Circle{}
	}
	return &Square{}
}

func main() {
	sh := ShapeFactory{}
	obj := sh.getShape("square")
	fmt.Println(obj.draw())

}
