package main

import (
	"fmt"
)

func main() {
	//unbuffered channel and buffered channel
	ch := make(chan int)
	defaultChannel := make(chan int)
	fmt.Printf("channel len is %d,cap is %d\n", len(defaultChannel), cap(defaultChannel))

	ch <- 1
	fmt.Println("go to the end ")
}
