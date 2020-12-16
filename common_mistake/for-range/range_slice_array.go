package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//test 1
	rangeSlice()

	//test 2
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
	}

	time.Sleep(4 * time.Second)
}

func rangeSlice() {
	//Using goroutines on loop iterator variables
	var values = []string{
		"a",
		"b",
		"c",
	}
	var wg sync.WaitGroup
	for _, val := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(val)
		}()
	}

	wg.Wait()
	fmt.Println("the program has exited")
}
