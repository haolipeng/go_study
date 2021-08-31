package main

import (
	"fmt"
	"testing"
)

//知识点1
//在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）作为函数参数
//这样的引用类型都是默认使用引用传递（即使没有显式的指出用指针传递）。

//知识点2
//命名的返回值,尽量使用命名返回值：会使代码更清晰、更简短、更易读
func min(first, second int) (sum int) {
	sum = first + second
	return
}

//知识点3
//传递指针不但可以节省内存，而且赋予了函数直接修改外部变量的能力

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

//知识点4
//函数类型作为函数参数
func add(a, b int) int {
	fmt.Println("call add function!")
	return a + b
}

func callback(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func TestFunc(t *testing.T) {
	res := min(1, 2)
	fmt.Printf("result:%d\n", res)

	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)

	//测试函数类型作为函数参数
	res = callback(10, 5, add)
	fmt.Println("result:", res)
}
