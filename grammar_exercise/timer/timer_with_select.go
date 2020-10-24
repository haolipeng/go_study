package main

import (
	"context"
	"fmt"
	"time"
)

//Ticker和Timer的区别是什么?
//Ticker是永久定时器，如过不调用Stop函数就会一直触发
//Timer是单次定时器，如果想定时器再次触发，则需调用Reset函数
//深入理解下time.After
/*
* 模拟生产者和消费者模型，对数据进行处理
 */
//数据生产者 写入数据到channel中20次
func dataProducer(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
}

//数据消费者,使用select从channel通道中读取数据，和定时器一起使用
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
		/********************重要的话说三遍******************/
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		///进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		//如果定时器没有到达定时时间，则gc不会启动垃圾回收
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

	//等待10秒，输出结果
	time.Sleep(10 * time.Second)
}

/*Tick是对NewTicker的封装*/
func useTick() {
	c := time.Tick(time.Second) //每次走一秒，输出日期
	for now := range c {
		fmt.Println("time at ", now)
	}
}

/*NewTimer 例子
Timer定时器是一次性定时器，想执行下次就需要手动Reset
*/
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
		case t := <-ticker.C: //输出当前时间
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

func main() {
	//useTick()
	//useNewTimer()
	//useNewTicker()

	//在select中使用time.After
	timeAfterInSelect()
	//ctx, _ := context.WithTimeout(context.Background(), time.Duration(8 * time.Second))
	//go useNewTimerWithContext(ctx)

	//time.Sleep(30 * time.Second)
}
