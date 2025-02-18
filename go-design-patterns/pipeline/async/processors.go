package async

import (
	"context"
	"log"
	"sync"
)

// SqProcessor 计算平方
type SqProcessor struct {
}

func (s *SqProcessor) Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) <-chan int {
	defer wg.Done()
	outChannel := make(chan int)

	go func() {
		defer close(outChannel)

		for {
			select {
			case s, ok := <-dataChan:
				if !ok {
					log.Println("sq data channel closed!")
					return
				}
				outChannel <- s * s
			case <-ctx.Done():
				log.Println("Sq processor received cancel signal")
				return
			}
		}
	}()

	return outChannel
}

// SumProcessor 累加
type SumProcessor struct {
}

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
				sum += s
				outChannel <- sum
			case <-ctx.Done():
				log.Println("Sum processor received cancel signal")
				return
			}
		}
	}()

	return outChannel
}
