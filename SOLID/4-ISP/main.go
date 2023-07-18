package main

// Interface Segreagation Principle
// You should not put too much into an interface

type Document struct {
}

// This interface have too much
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFuncPrinter struct {
}

func (mp *MultiFuncPrinter) Print(d Document) {

}
func (mp *MultiFuncPrinter) Fax(d Document) {

}
func (mp *MultiFuncPrinter) Scan(d Document) {

}

// The interface is not appropiate for an old printer
type OldFuncPrinter struct {
}

func (mp *OldFuncPrinter) Print(d Document) {
	// ok

}
func (mp *OldFuncPrinter) Fax(d Document) {
	//  does not support
	panic("not supported")

}
func (mp *OldFuncPrinter) Scan(d Document) {
	//  does not support
	panic("not supported")
}

// ISP: seggregate into multiple interfaces

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// only prints
type MyPrinter struct {
}

func (mp MyPrinter) Print(d Document) {

}

//  prints and scans
type Photocopier struct {
}

func (ph Photocopier) Print(d Document) {

}

func (ph Photocopier) Scan(d Document) {

}

// Combine interfaces
type MultiFuncDevice interface {
	Printer
	Scanner
}

// decorator
type MultiFuncMachine struct {
	printer Printer
	scanner Scanner
}

func (ph MultiFuncMachine) Print(d Document) {
	ph.printer.Print(d)
}

func (ph MultiFuncMachine) Scan(d Document) {
	ph.scanner.Scan(d)
}

func main() {
	doc := Document{}
	ofp := OldFuncPrinter{}
	// ok
	ofp.Print(doc)
	// fails
	//ofp.Scan(doc)

}
