package main

import (
	"time"
	"fmt"
)

//var chanElemCnt int = 10
//var globalChannel = make(chan int)

//signChannel用于保证main函数等待go协程执行完毕
//var signChannel = make(chan int, 2)

//close一个通道会唤醒所有等待在通道上的goroutine协程
//这些writer协程发现该channel已经是closed状态，就panic了
//使用 comma, ok 语法来区分channel中返回的是零值还是buffer值
func closeChannel() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)

	go func() {
		close(ch)
	}()
	time.Sleep(2 * time.Second)

	x, ok := <-ch
	fmt.Println(x, ok)
}

//关闭一个空channel会触发panic
func closeNilChannel() {
	var a chan int
	close(a)
}

//一个sender，多个receiver，由sender来关闭channel，通知数据已经发送完毕
//一旦sender有多个，可能就无法判断数据是否完毕了，可以借助而外channel来做信号广播
//优雅的关闭channel通道，链接地址：https://go101.org/article/channel-closing.html

func main() {
	//panic: send on closed channel
	//closeChannel()

	//panic: close of nil channel
	closeNilChannel()
}
