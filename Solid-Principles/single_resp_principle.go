package main

// Single Responsibility Principle (SRP)
// Definition: A class should have only one reason to change, meaning it should have only one job or responsibility.

//******** Before SRP: ******//
/*
package main

import (
	"fmt"
)

type Order struct {
	ID    int
	Items []string
	Total float64
}

func (o Order) PrintInvoice() {
	fmt.Println("Invoice for Order ID:", o.ID)
	for _, item := range o.Items {
		fmt.Println("-", item)
	}
	fmt.Println("Total:", o.Total)
}

func main() {
	order := Order{ID: 1, Items: []string{"Shoes", "Shirt"}, Total: 2500}
	order.PrintInvoice()
}

*/

//**** After SRP
/*
type Order struct {
	ID    int
	Items []string
	Total float64
}
type InvoicePrinter struct{}

func (ip InvoicePrinter) Print(order Order) {
	fmt.Println("Invoice for Order ID:", order.ID)
	for _, item := range order.Items {
		fmt.Println("-", item)
	}
	fmt.Println("Total:", order.Total)
}

func main() {
	order := Order{ID: 1, Items: []string{"Shoes", "Shirt"}, Total: 2500}
	printer := InvoicePrinter{}
	printer.Print(order)
}
*/
