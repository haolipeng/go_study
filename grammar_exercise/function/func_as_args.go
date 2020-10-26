package main

import "fmt"

//函数类型作为函数参数
func add(a, b int) int {
	fmt.Println("call add function!")
	return a + b
}

func callback(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func main() {
	res := callback(10, 5, add)
	fmt.Println("result:", res)
}
