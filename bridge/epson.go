package bridge

import "fmt"

type Epson struct {
}

func NewEpson() Printer {
	return &Epson{}
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
