package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func NewMac(printer Printer) Computer {
	return &Mac{
		printer: printer,
	}
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
