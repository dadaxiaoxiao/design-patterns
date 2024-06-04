package adapter

import (
	"fmt"
	"testing"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

func Test(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := NewWindowsAdapter(windowsMachine)

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
