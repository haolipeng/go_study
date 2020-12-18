package main

import "fmt"

func main() {
	//创建多维map
	m := make(map[string]map[string]string)

	m1 := make(map[string]string)
	m1["b"] = "c"
	m["a"] = m1
	fmt.Printf("map content:%v\n", m)
	fmt.Println("you are the best")

	//添加元素时，也要注意
	if _, ok := m["a"]; ok {
		m["a"]["c"] = "d"
	} else {
		//每次添加数据时，都要重新创建一个map结构
		m2 := make(map[string]string)
		m2["c"] = "d"
		m["a"] = m2
	}
	fmt.Printf("map content:%v\n", m)
}
