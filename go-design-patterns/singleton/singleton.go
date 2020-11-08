package main

import (
	"fmt"
	"sync"
	"unsafe"
)

//演示单例模式，采用sync.Once特性

var once sync.Once
var SingletonInstance *Singleton

type Singleton struct {
	name string
}

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singleton obj.")
		SingletonInstance = new(Singleton)
	})

	return SingletonInstance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("pionter:%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}

	wg.Wait()
}
