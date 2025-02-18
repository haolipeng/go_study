package async

import (
	"context"
	"log"
	"sync"
)

// SqProcessor 平方处理器
// 对输入的数据计算平方值
type SqProcessor struct {
}

// Process 实现处理器接口
// 从输入通道读取数据，计算平方后发送到输出通道
// 当输入通道关闭或收到取消信号时退出
func (s *SqProcessor) Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) <-chan int {
	defer wg.Done()
	outChannel := make(chan int)

	go func() {
		defer close(outChannel)
		
		for {
			select {
			// 从输入通道读取数据
			case s, ok := <-dataChan:
				if !ok {
					log.Println("sq data channel closed!")
					return
				}
				// 处理数据并确保发送成功
				result := s * s
				select {
				case outChannel <- result:
					// 数据发送成功
				case <-ctx.Done():
					// 收到取消信号，但仍要确保当前数据发送出去
					log.Println("Sq processor received cancel signal, sending last data")
					outChannel <- result
					return
				}
				
			case <-ctx.Done():
				// 继续处理输入通道中的剩余数据
				log.Println("Sq processor draining remaining data")
				for s := range dataChan {
					outChannel <- s * s
				}
				log.Println("Sq processor completed")
				return
			}
		}
	}()

	return outChannel
}

// SumProcessor 累加处理器
// 对输入的数据进行累加计算
type SumProcessor struct {
}

// Process 实现处理器接口
// 从输入通道读取数据，与之前的累加和相加后发送到输出通道
// 当输入通道关闭或收到取消信号时退出
func (s *SumProcessor) Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) <-chan int {
	defer wg.Done()
	outChannel := make(chan int)

	go func() {
		defer close(outChannel)
		var sum = 0

		for {
			select {
			case s, ok := <-dataChan:
				if !ok {
					log.Println("sum data channel closed!")
					return
				}
				// 处理数据并确保发送成功
				sum += s
				select {
				case outChannel <- sum:
					// 数据发送成功
				case <-ctx.Done():
					// 收到取消信号，但仍要确保当前数据发送出去
					log.Println("Sum processor received cancel signal, sending last data")
					outChannel <- sum
					return
				}

			case <-ctx.Done():
				// 继续处理输入通道中的剩余数据
				log.Println("Sum processor draining remaining data")
				for s := range dataChan {
					sum += s
					outChannel <- sum
				}
				log.Println("Sum processor completed")
				return
			}
		}
	}()

	return outChannel
}
