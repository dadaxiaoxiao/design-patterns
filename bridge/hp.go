package bridge

import "fmt"

type Hp struct {
}

func NewHp() Printer {
	return &Hp{}
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
