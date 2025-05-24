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
