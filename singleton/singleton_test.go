package singleton

import (
	"fmt"
	"testing"
)

func TestBaseSingleton(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			getBaseInstance()
		}()
	}
	fmt.Scanln()
}

func TestSingleton(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			getInstance()
		}()
	}
	fmt.Scanln()
}
