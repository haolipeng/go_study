package async

import (
	"context"
	"sync"
)

// IProcessor 定义数据处理器接口
type IProcessor interface {
	Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) <-chan int
}

// ISource 定义数据源接口
type ISource interface {
	Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error) <-chan int
}

// ISink 定义数据接收器接口
type ISink interface {
	Process(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error)
}

// IError 定义错误处理器接口
type IError interface {
	Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancel context.CancelFunc)
}
