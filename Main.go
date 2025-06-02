package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("hello")
}

func main() {
	var obj Shape = &Circle{}
	obj.draw()

}
