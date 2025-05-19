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
