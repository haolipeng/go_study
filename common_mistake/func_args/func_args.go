package main

import "fmt"

//数组作为函数参数,golang和C/C++不同，golang向函数传递数组时，函数会得到原始数组数据的一份拷贝
func FuncArgsIsArray(arr [3]int) {
	fmt.Printf("in function array pointer:%p\n", &arr)
	arr[0] = 10
	fmt.Println("in function modify, array:", arr)
}

//切片作为函数参数
func FuncArgsIsSlice(s []int) {
	fmt.Printf("in function slice pointer:%p\n", s)
	s[0] = 10
	fmt.Println("in function modify, slice:", s)
}

//向函数中传递数组时，函数会得到原始数组数据的一份拷贝，如果在函数内部想更新数据，是不可能的
func main() {
	fmt.Println("--------------测试数组作为函数参数--------------")
	x := [3]int{1, 2, 3}
	fmt.Printf("before function array pointer:%p\n", &x)
	FuncArgsIsArray(x)
	fmt.Println("after array modify, array:", x)

	fmt.Println("--------------测试切片作为函数参数--------------")
	y := []int{1, 2, 3}
	fmt.Printf("before function slice pointer:%p\n", y)
	FuncArgsIsSlice(y)
	fmt.Println("after slice modify, slice:", y)
}
