package async

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// ConsoleSink 控制台输出接收器
// 将接收到的数据打印到控制台
type ConsoleSink struct {
}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{}
}

// Process 实现接收器接口
// 从输入通道读取数据并打印到控制台
// 当输入通道关闭或收到取消信号时退出
func (s *ConsoleSink) Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) {
	go func() {
		defer wg.Done()

		for {
			select {
			case val, ok := <-dataChan:
				if !ok {
					log.Println("sink data channel closed!")
					return
				}
				fmt.Printf("sink value: %v\n", val)

			case <-ctx.Done():
				// 继续处理输入通道中的剩余数据
				log.Println("Sink draining remaining data")
				for val := range dataChan {
					fmt.Printf("sink value (draining): %v\n", val)
				}
				log.Println("Sink completed")
				return
			}
		}
	}()
}
