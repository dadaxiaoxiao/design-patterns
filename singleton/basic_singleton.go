package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

// baseSingleton 基础单例
// 通过 sync.Mutex 实现线程安全
type baseSingleton struct {
}

var baseSingletonInstance *baseSingleton

func getBaseInstance() *baseSingleton {
	if baseSingletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if baseSingletonInstance == nil {
			baseSingletonInstance = &baseSingleton{}
			fmt.Println("Creating single instance now")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return baseSingletonInstance
}
