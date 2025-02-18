package async

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// ConsoleSink 输出到命令行
type ConsoleSink struct {
}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{}
}

func (s *ConsoleSink) Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) {

	go func() {
		defer wg.Done()

		for {
			select {
			case val, ok := <-dataChan:
				if ok {
					fmt.Printf("sink value: %v\n", val)
				} else {
					log.Println("sink data channel closed!")
					return
				}
			case <-ctx.Done():
				log.Println("Sink received cancel signal")
				return
			}
		}
	}()
}
