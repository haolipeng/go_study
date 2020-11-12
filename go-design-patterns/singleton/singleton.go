package singleton

import (
	"fmt"
	"sync"
)

//演示单例模式，采用sync.Once特性

var once sync.Once
var Instance *Singleton

type Singleton struct {
	name string
}

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singleton obj.")
		Instance = new(Singleton)
	})

	return Instance
}
