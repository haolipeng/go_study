package main

import (
	"fmt"
	"testing"
	"time"
)

//模拟生产者和消费者模型，对数据进行处理

//数据生产者 写入数据到channel中
func dataProducer(ch chan int) {
	for i := 0; i < 40; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 1000)
	}
}

//数据消费者,使用select从channel通道中读取数据，和定时器一起使用
func dataConsumer(ch chan int) {
	duration := 2 * time.Second

	//NewTime + Reset = NewTicker
	tm := time.NewTimer(duration)
	tick := time.NewTicker(duration)
	for {
		select {
		case data := <-ch: //select会"随机"选择data1分支或data2分支来执行。
			fmt.Println("data1 case")
			fmt.Println(data)
			/*		case data := <-ch:
					fmt.Println("data2 case")
					fmt.Println(data)*/
		case <-tm.C:
			fmt.Println("call timer in 1s second")
			tm.Reset(time.Second)
		case <-tick.C:
			fmt.Println("call ticker in 1s second")
		//如果ch通道中持续有数据，则不会触发 <-time.After(time.Second) 此case语句
		/********************重要的话说三遍******************/
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		//每进行一次select选择时，time.After会重新创建一个定时器，而以前的定时器还遗留在时间堆中等待gc释放
		//如果定时器没有到达定时时间，则gc不会启动垃圾回收
		case <-time.After(duration):
			fmt.Println("timeout")
			tm.Stop()
			tick.Stop()
		}
	}
}

//在select多路选择器中使用time.After是错误的，After可能不触发
func TestTimeAfterInSelect(t *testing.T) {
	ch := make(chan int, 1)
	go dataProducer(ch)
	go dataConsumer(ch)

	//等待10秒，输出结果
	time.Sleep(10 * time.Second)
	t.Log("program is exited")
}
