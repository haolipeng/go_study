package main

import (
	"fmt"
	"go_study/grammar_exercise/init/pack"
	"go_study/grammar_exercise/init/util"
)

//1.init函数先于main函数自动执行，不能被其他函数调用
//2.init函数没有输入参数、返回值
//3.每个包可以有多个init函数
//4.包的每个源文件也可以有多个init函数，这点比较特殊
//5.同一个包的init执行顺序，golang并没有明确定义，编程时不要依赖这个执行顺序
//6.不同包的init函数按照包导入的依赖关系决定执行顺序

var T int64 = a()

func a() int64 {
	fmt.Println("calling a()")
	return 2
}

//一个源文件中可以有多个init函数，有这个必要吗？
func init() {
	fmt.Println("I am init() function")
}

func init() {
	fmt.Println("I am init() function Too Too,haha")
}

func main() {
	//结论:变量初始化 -> init() -> main()
	//init() //error 包内的init函数不可以被显式的调用

	//不同包的init函数按照包导入的依赖关系决定执行顺序,import列表中先导入的pack包
	fmt.Printf("util value is: %d\n", util.UtilVar)
	fmt.Printf("pack value is: %d\n", pack.PackVar)

	fmt.Println("the program has exited")
}
