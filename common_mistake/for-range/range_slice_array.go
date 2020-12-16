package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	//test 0
	fmt.Println("-----------------test1----------------------")
	testSlice()
	time.Sleep(3 * time.Second)

	fmt.Println("-----------------test2----------------------")
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three

	fmt.Println("-----------------test3----------------------")
	data2 := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data2 {
		go v.print()
	}
	//goroutines print: three, three, three

	/*
		test 1
		x := []string{"a", "b", "c"}

		//遍历slice时，返回的第一个值是index索引
		for _, v := range x {
			fmt.Printf("out of goroutine -> value:%v pointer:%v\n", v, &v)
		}
		fmt.Println("-------------------------------------------------")
		for _, v := range x {
			go func() {
				fmt.Printf("in goroutine -> value:%v pointer:%v\n", v, &v)
			}()
			//time.Sleep(time.Second)
		}*/

	time.Sleep(4 * time.Second)
}

//1.选择将变量作为函数的参数传递
//2.将迭代的变量值保存到一个临时变量中
func testSlice() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	//goroutines print: three, three, three
}
