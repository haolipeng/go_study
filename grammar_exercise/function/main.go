package main

import "fmt"

//知识点1
//在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）
//这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。

//知识点2
//命名的返回值,尽量使用命名返回值：会使代码更清晰、更简短、更易读
func min(first, second int) (sum int) {
	sum = first + second
	return
}

//知识点3
//传递指针不但可以节省内存，而且赋予了函数直接修改外部变量的能力
func Mutiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	res := min(1, 2)
	fmt.Printf("result:%d\n", res)

	//test mutiply
	n := 0
	reply := &n
	Mutiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)
}