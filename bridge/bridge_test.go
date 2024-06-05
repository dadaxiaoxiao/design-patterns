package bridge

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	hpPrinter := NewHp()
	epsonPrinter := NewEpson()

	macComputer := NewMac(nil)
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := NewWindows(nil)
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()

}
