package async

import (
	"context"
	"log"
	"sync"
)

// ErrorPolicyExit 错误处理
type ErrorPolicyExit struct {
}

func NewErrorPolicyExit() *ErrorPolicyExit {
	return &ErrorPolicyExit{}
}

func (p *ErrorPolicyExit) Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancel context.CancelFunc) {
	for {
		select {
		case err, ok := <-errChan:
			if !ok {
				log.Println("error channel closed and exit!")
				return
			}

			log.Printf("Receive error %v\n", err)
			cancel()
		}
	}
}
