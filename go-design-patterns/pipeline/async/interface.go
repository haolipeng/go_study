package async

import (
	"context"
	"sync"
)

type IProcessor interface {
	Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) <-chan int
}

type ISource interface {
	Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error) <-chan int
}

type ISink interface {
	Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error)
}

type IError interface {
	Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancel context.CancelFunc)
}
