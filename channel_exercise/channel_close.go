package main

import (
	"time"
	"fmt"
)

//close一个通道会唤醒所有等待在通道上的goroutine协程
//这些writer协程发现该channel已经是closed状态，就panic了
//使用 comma, ok 语法来区分channel中返回的是零值还是buffer值
func closeChannel() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)

	/*
	go func() {
		close(ch)
	}()
	*/

	time.Sleep(2 * time.Second)

	//comma语法
	x, ok := <-ch
	fmt.Println(x, ok)
}

//测试1：关闭空channel会触发panic
func closeNilChannel() {
	var a chan int
	close(a)
}

func main() {
	//panic: send on closed channel
	closeChannel()

	//panic: close of nil channel
	///closeNilChannel()
}
