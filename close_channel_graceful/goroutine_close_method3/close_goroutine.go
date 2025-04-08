package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func somework(ctx context.Context, wg *sync.WaitGroup, stop <-chan struct{}) chan int {
	defer wg.Done()
	out := make(chan int)

	go func() {
		defer close(out) // 关闭下游数据通道
		pktID := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[somework] 收到context cancel 信号")
				return
			case <-stop:
				fmt.Println("[somework] 收到stop channel 停止信号")
				return
			default:
				// 模拟收包
				fmt.Printf("[somework] 收到数据包 %d\n", pktID)
				out <- pktID
				pktID++
				time.Sleep(500 * time.Millisecond)
				if pktID >= 10 {
					fmt.Println("[somework] 收到10个数据包，正常退出")
					return
				}
			}
		}
	}()

	return out
}

func main() {
	//1.create context with cancel
	cancelCtx, _ := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)

	stop := make(chan struct{})
	out := somework(cancelCtx, &wg, stop)

	go func() {
		for {
			v, ok := <-out
			if !ok {
				fmt.Println("out channel is closed")
				break
			}
			fmt.Println("value:", v)
		}
	}()

	time.Sleep(3 * time.Second)
	//有三种方式关闭go协程
	//close(stop) // 关闭stop channel
	//cancel() // 外部关闭context
	//休眠几秒，让程序有序退出
	time.Sleep(time.Second * 3)

	wg.Wait()
}
