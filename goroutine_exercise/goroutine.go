package main

import (
	"runtime"
	"fmt"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		//Gosched函数是将cpu时间片让出，通过通信来共享是golang的特色
		runtime.Gosched()
		fmt.Println(s)
	}
}

func channelPractise() {
	ci := make(chan int, 2)

	ci <- 10
	ci <- 20
	close(ci)
	/*
	first := <-ci
	second := <-ci
	fmt.Printf("first is %d,second is %d\r\n",first,second)
	*/

	//使用for range来遍历channel中元素
	for value := range ci {
		fmt.Printf("value is %d\n", value)
	}
}

func main() {
	go say("hello world")
	say("hello")

	channelPractise()
}
