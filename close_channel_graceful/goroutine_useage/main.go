package main

import (
	"context"
	"fmt"
	"time"
)

func somework(stop chan struct{}, ctx context.Context) chan int {
	out := make(chan int)

	go func() {
		defer close(out) // 在协程关闭前,关闭下游数据通道
		pktID := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context cancel function called,exit")
				return
			case <-stop:
				fmt.Println("stop channel receive signal,exit")
				return
			case out <- pktID:
				// 模拟收包
				fmt.Printf("[somework] 构造数据包 %d\n", pktID)
				pktID++
				time.Sleep(500 * time.Millisecond)
				if pktID >= 10 {
					fmt.Println("[somework] 构造10个数据包，正常退出")
					return
				}
			}
		}
	}()

	return out
}

func main() {
	stop := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	//produce data
	out := somework(stop, ctx)

	//消费数据
	go func() {
		for {
			v, ok := <-out
			//如channel通道已关闭，则直接退出
			if !ok {
				fmt.Println("out channel is closed")
				break
			}
			fmt.Println("value:", v)
		}
	}()

	time.Sleep(time.Second * 2)
	//有三种方式关闭go协程
	//close(stop) // 关闭stop channel
	cancel() // 外部关闭context
	//休眠几秒，让程序有序退出
	time.Sleep(time.Second * 3)
}
