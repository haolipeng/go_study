package sync

import (
	"context"
	"log"
	"sync"
)

// 莎士比亚的十四行诗
var LINES = `Shakespeare Sonnet 12
When I do count the clock that tells the time
And see the brave day sunk in hideous night
When I behold the violet past prime
And sable curls all silver'd o'er with white
When lofty trees I see barren of leaves
Which erst from heat did canopy the herd
And summer's green, all girded up in sheaves
Born on the bier with white and bristly beard
Then of thy beauty do I question make
That thou among the wastes of time must go
Since sweets and beauties do themselves forsake
And die as fast as they see others grow
And nothing 'gainst Time's scythe can make defence
Save breed, to brave him when he takes thee hence`

type ProcessorManager struct {
	source ISource      //一个数据源
	sink   ISink        //一个输出
	ps     []IProcessor //多个处理器
}

func NewProcessorManager() *ProcessorManager {
	return &ProcessorManager{}
}

// AddProcessor add processor
func (pm *ProcessorManager) AddProcessor(processor IProcessor) {
	pm.ps = append(pm.ps, processor)
}

// AddSource add source
func (pm *ProcessorManager) AddSource(source ISource) {
	pm.source = source
}

// AddSink add sink
func (pm *ProcessorManager) AddSink(sink ISink) {
	pm.sink = sink
}

// Run 执行整个处理管道，包含数据源读取、处理器链处理、数据落地三个阶段
// 单个并发度
func (pm *ProcessorManager) Run(ctx context.Context) error {
	var err error

	// 第一阶段：从数据源获取数据流
	// in 为一个channel，用于异步接收数据
	in, err := pm.source.Process(ctx)
	if err != nil {
		return err
	}

	// 第二阶段：遍历数据流，执行处理管道
	for data := range in {
		// 依次执行所有处理器
		// 每个处理器的输出作为下一个处理器的输入
		for _, p := range pm.ps {
			data, err = p.Process(ctx, data)
			if err != nil {
				log.Printf("process err %s\n", err)
				//不返回错误，直接处理下一条数据
				//return err
			}
		}

		// 第三阶段：数据落地
		// 将处理后的数据交给sink处理器进行持久化
		err = pm.sink.Process(ctx, data)
		if err != nil {
			log.Printf("Sink err %s\n", err)
			return err
		}
	}
	return nil
}

// RunN 并发执行处理管道，支持多个goroutine同时处理数据
// 多个并发度
func (pm *ProcessorManager) RunN(ctx context.Context, maxCnt int) error {
	// 第一阶段：初始化数据源
	// 获取数据流channel
	in, err := pm.source.Process(ctx)
	if err != nil {
		return err
	}

	// 定义单条数据的处理函数
	// 包含处理器链执行和数据落地两个步骤
	syncProcess := func(data any) {
		// 依次执行所有处理器
		// 处理器的输出作为下一个处理器的输入
		for _, v := range pm.ps {
			data, err = v.Process(ctx, data)
			if err != nil {
				log.Printf("process err %s\n", err)
				return
			}
		}

		// 数据落地操作
		err := pm.sink.Process(ctx, data)
		if err != nil {
			log.Printf("sink err %s\n", err)
			return
		}
	}

	// 使用WaitGroup控制并发goroutine的生命周期
	wg := sync.WaitGroup{}
	wg.Add(maxCnt)

	// 启动多个goroutine并发处理数据
	// 每个goroutine独立消费数据源channel
	for i := 0; i < maxCnt; i++ {
		go func() {
			defer wg.Done() // 确保goroutine退出时减少计数

			// 循环处理数据，直到channel关闭
			for data := range in {
				syncProcess(data)
			}
		}()
	}

	// 等待所有goroutine完成处理
	wg.Wait()
	return nil
}
