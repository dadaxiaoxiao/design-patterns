package adapter

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func NewWindowsAdapter(windowMachine *Windows) Computer {
	return &WindowsAdapter{windowMachine: windowMachine}
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
