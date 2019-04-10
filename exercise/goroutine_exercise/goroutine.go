package main

import (
	"fmt"
	"runtime"
)

//轻量级的线程
//非抢占式多任务处理，由协程主动交出控制权
//传统的线程是抢占式多任务处理，线程的执行由操作系统来调度，有时语句执行一半被迫停止
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
