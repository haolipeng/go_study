package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建上下文和停止信号通道
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 数据通道
	captureToParse := make(chan int)
	parseToStore := make(chan int)

	// 停止信号通道
	stopCapture := make(chan struct{})
	stopParse := make(chan struct{})
	stopStore := make(chan struct{})

	var wg sync.WaitGroup

	// 启动协程（存储 → 解析 → 收包）
	wg.Add(3)
	go store(ctx, &wg, parseToStore, stopStore)
	go parse(ctx, &wg, captureToParse, parseToStore, stopParse, stopStore)
	go capture(ctx, &wg, captureToParse, stopCapture, stopParse)

	// 运行一段时间后触发停止
	time.Sleep(3 * time.Second)
	fmt.Println("\n===> 开始有序停止流程")

	// 触发停止（从收包开始）
	close(stopCapture)

	// 等待所有协程退出
	time.Sleep(10 * time.Second)
	wg.Wait()
	fmt.Println("所有协程已安全退出")
}

// 收包阶段
func capture(ctx context.Context, wg *sync.WaitGroup, out chan<- int, stopCapture <-chan struct{}, stopParse chan<- struct{}) {
	defer wg.Done()
	defer close(out)       // 关闭下游数据通道
	defer close(stopParse) // 触发解析阶段停止

	pktID := 1
	for {
		select {
		case <-ctx.Done():
			return
		case <-stopCapture:
			fmt.Println("[capture] 收到停止信号")
			return
		default:
			// 模拟收包
			fmt.Printf("[capture] 收到数据包 %d\n", pktID)
			out <- pktID
			pktID++
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

// 解析阶段
func parse(ctx context.Context, wg *sync.WaitGroup, in <-chan int, out chan<- int, stopParse <-chan struct{}, stopStore chan<- struct{}) {
	defer wg.Done()
	defer close(out)       // 关闭下游数据通道
	defer close(stopStore) // 触发存储阶段停止

	for {
		select {
		case <-ctx.Done():
			return
		case <-stopParse:
			fmt.Println("[parse] 收到停止信号")
			return
		case pktID, ok := <-in:
			if !ok {
				return
			}
			// 模拟解析
			fmt.Printf("[parse] 解析数据包 %d\n", pktID)
			time.Sleep(1000 * time.Millisecond)
			out <- pktID
		}
	}
}

// 存储阶段
func store(ctx context.Context, wg *sync.WaitGroup, in <-chan int, stopStore <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): //context超时或取消操作
			return
		case <-stopStore:
			fmt.Println("[store] 收到停止信号")
			return
		case pktID, ok := <-in:
			if !ok {
				return
			}
			// 模拟存储
			fmt.Printf("[store] 存储数据包 %d\n", pktID)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
