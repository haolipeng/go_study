package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	defaultChannel := make(chan int)
	fmt.Printf("channel len is %d,cap is %d\n", len(defaultChannel), cap(defaultChannel))

	fmt.Println("hello world")
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	ch <- 1
	fmt.Println("go to the end ")
}
