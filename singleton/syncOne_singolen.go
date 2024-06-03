package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type Singleton struct {
}

var singleton *Singleton

func getInstance() *Singleton {
	if singleton == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now")
			singleton = &Singleton{}
		})
	} else {
		fmt.Println("Single instance already created")
	}
	return singleton
}
