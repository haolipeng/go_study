package async

import (
	"context"
	"sync"
)

type ProcessorManager struct {
	source  ISource
	sink    ISink
	err     IError
	ps      []IProcessor
	errChan chan error
}

func NewProcessorManager() *ProcessorManager {
	return &ProcessorManager{errChan: make(chan error, 2)}
}

func (m *ProcessorManager) AddProcessor(processor IProcessor) {
	m.ps = append(m.ps, processor)
}

func (m *ProcessorManager) AddSource(source ISource) {
	m.source = source
}

func (m *ProcessorManager) AddSink(sink ISink) {
	m.sink = sink
}

func (m *ProcessorManager) AddError(err IError) {
	m.err = err
}

func (m *ProcessorManager) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg = sync.WaitGroup{}

	// 组装pipeline,如何组装这个pipeline的呢
	// 1. 先从source获取数据
	wg.Add(1)
	dataChan := m.source.Process(ctx, &wg, m.errChan)

	// 2. 将数据传递给processor进行处理
	for _, v := range m.ps {
		wg.Add(1)
		dataChan = v.Process(ctx, &wg, dataChan, m.errChan)
	}

	// 3. 将处理后的数据传递给sink进行处理
	wg.Add(1)
	m.sink.Process(ctx, &wg, dataChan, m.errChan)

	go func() {
		wg.Wait()
		//所有处理完成后，关闭err错误通道
		close(m.errChan)
	}()

	// 错误通道内部逻辑是for循环，阻塞，错误处理集中处理
	// 出现错误则通知退出，可灵活定制处理策略
	m.err.Process(ctx, &wg, m.errChan, cancel)
}
