package main

/*
import "fmt"

////  ****** Before interface segeration ********** //

import "fmt"

type Printer interface {
	PrintDocument()
	ScanDocument()
	FaxDocument()
}

type SimplePrinter struct{}

func (sp SimplePrinter) PrintDocument() {
	fmt.Println("Printing document...")
}

func (sp SimplePrinter) ScanDocument() {
	// Not implemented
}

func (sp SimplePrinter) FaxDocument() {
	// Not implemented
}

func main() {
	var printer Printer = SimplePrinter{}
	printer.PrintDocument()
}


//  ****** After interface segeration ********** //

// Segregated interfaces
type Printer interface {
	PrintDocument()
}

type Scanner interface {
	ScanDocument()
}

type Faxer interface {
	FaxDocument()
}

// Only implements what it needs
type SimplePrinter struct{}

func (sp SimplePrinter) PrintDocument() {
	fmt.Println("Printing document...")
}

// Another printer that does all
type AdvancedPrinter struct{}

func (ap AdvancedPrinter) PrintDocument() {
	fmt.Println("Advanced: Printing document...")
}

func (ap AdvancedPrinter) ScanDocument() {
	fmt.Println("Advanced: Scanning document...")
}

func (ap AdvancedPrinter) FaxDocument() {
	fmt.Println("Advanced: Faxing document...")
}

func main() {
	var p Printer = SimplePrinter{}
	p.PrintDocument()

	var ap Printer = AdvancedPrinter{}
	ap.PrintDocument()
}
*/
