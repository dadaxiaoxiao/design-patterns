package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func NewWindows(printer Printer) Computer {
	return &Windows{
		printer: printer,
	}
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
