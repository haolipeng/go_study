package main

import "fmt"

func main() {
	//var x string = nil
	//nil is a predeclared identifier representing the zero value for
	//a pointer, channel, func, interface, map, or slice type.
	//Type must be a pointer, channel, func, interface, map, or slice type

	var x string //默认值为空字符串""
	if x == "" {
		fmt.Println("x string is default empty string")
	}

	fmt.Println("string can't be nil")
}
