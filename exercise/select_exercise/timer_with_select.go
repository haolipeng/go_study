package main

import (
	"fmt"
	"time"
)

//Ticker和Timer的区别是什么?
//Ticker是永久定时器，如过不调用Stop函数就会一直触发
//Timer是单次定时器

func dataProducer(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
}

func dataConsumer(ch chan int) {
	tm := time.NewTimer(time.Second)
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-tm.C:
			fmt.Println("call new timer!")
			tm.Reset(time.Second)
		//如果ch通道中持续有数据，则不会触发此case语句
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}
}

/*在select多路选择器中使用time.After是错误的
 */
func testTimerInSelect() {
	ch := make(chan int, 1)
	go dataProducer(ch)
	go dataConsumer(ch)
}

/*Tick是对NewTicker的封装*/
func TickUse() {
	c := time.Tick(time.Second)
	for now := range c {
		fmt.Println("time at ", now)
	}
}

/*NewTimer 例子*/
func UseNewTimer() {
	onceTimer := time.NewTimer(time.Second)
	everyTimer := time.NewTimer(time.Second * 2)

	for {
		select {
		case <-onceTimer.C:
			fmt.Println("call 1s once timer!")
		case <-everyTimer.C:
			fmt.Println("call 2s every timer")
			everyTimer.Reset(time.Second * 2)
		}
	}
}

/*NewTicker 例子*/
func UseNewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("get Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time:", t)
		}
	}
}

func main() {

	//TickUse()
	//UseNewTicker()
	UseNewTimer()
}
