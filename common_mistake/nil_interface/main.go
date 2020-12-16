package main

import "fmt"

//interface虽然看起来像指针，但并不是指针。interface变量仅在类型和值为“nil”时才为“nil”。
//interface的类型和值会根据用于创建对应interface变量的类型和值的变化而变化。
func main() {
	var data *byte
	var in interface{}
	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true

	in = data
	fmt.Println(in, in == nil)
}
