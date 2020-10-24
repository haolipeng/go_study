package main

import (
	"context"
	"fmt"
	"time"
)

//Ticker和Timer的区别是什么?
//Ticker是永久定时器，如过不调用Stop函数就会一直触发
//Timer是单次定时器，如果想定时器再次触发，则需调用Reset函数

//数据生产者 从channel通道中读取数据20次
func dataProducer(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
}

//数据消费者,使用多路选择器从channel通道中读取数据，和定时器一起使用
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
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}
}

//在select多路选择器中使用time.After是错误的，After可能不触发
func timeAfterInSelect() {
	ch := make(chan int, 1)
	go dataProducer(ch)
	go dataConsumer(ch)
}

/*Tick是对NewTicker的封装*/
func useTick() {
	c := time.Tick(time.Second)
	for now := range c {
		fmt.Println("time at ", now)
	}
}

/*NewTimer 例子*/
func useNewTimer() {
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
func useNewTicker() {
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

func useNewTimerWithContext(ctx context.Context) {
	onceTimer := time.NewTimer(time.Second)
	everyTimer := time.NewTimer(time.Second * 2)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("useNewTimerWithContext function return!")
			return
		case <-onceTimer.C:
			fmt.Println("call 1s once timer!")
		case <-everyTimer.C:
			fmt.Println("call 2s every timer")
			everyTimer.Reset(time.Second * 2)
		}
	}
}

/*func main() {
	//useTick()
	//useNewTimer()
	//useNewTicker()
	//timeAfterInSelect()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(8 * time.Second))
	go useNewTimerWithContext(ctx)

	time.Sleep(30 * time.Second)
}
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
