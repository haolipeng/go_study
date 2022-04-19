package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	//1.写入
	m.Store("haolipeng", 33)
	m.Store("zhouyang", 34)

	//2.读取
	age, ok := m.Load("haolipeng")
	if ok {
		fmt.Println("haolipeng age:", age.(int))
	}

	//3.遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})
	//4.删除
	m.Delete("haolipeng")

	age, ok = m.Load("haolipeng")
	if ok {
		fmt.Println("haolipeng age:", age.(int))
	} else {
		fmt.Println("no haolipeng")
	}
}
