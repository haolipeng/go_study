package main

import "fmt"

func array_func_args(arr [3]int) {
	arr[0] = 10
	fmt.Println("after array modify, array:", arr)
}

func slice_func_args(arr []int) {
	arr[0] = 100
	fmt.Println("after slice modify, array:", arr)
}

//向函数中传递数组时，函数会得到原始数组数据的一份拷贝，如果在函数内部想更新数据，是不可能的
func main() {
	//数组
	x := [3]int{1, 2, 3}
	array_func_args(x)
	fmt.Println(x)

	fmt.Println("----------------------------")
	//slice
	y := []int{1, 2, 3}
	slice_func_args(y)
	fmt.Println(y)
}
