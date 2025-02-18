package functional_sync

import (
	"context"
	"log"
	"sync"
)

// LINES 莎士比亚的十四行诗
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

// PipeSourceFunc 数据源的函数类型定义
type PipeSourceFunc func(ctx context.Context) (chan any, error)
type PipeProcessFunc func(ctx context.Context, params any) (any, error)

func PipeProcessBuildAndRun(ctx context.Context, input PipeSourceFunc, funcs ...PipeProcessFunc) {
	var err error = nil

	// 输入
	dataChan, err := input(ctx)
	if err != nil {
		log.Printf("source error-%v\n", err.Error())
		return
	}

	// pipeline构建和执行
	for data := range dataChan {
		// 依次执行函数，上一个函数返回结果当做下个函数的输入参数
		for _, processFunc := range funcs {
			data, err = processFunc(ctx, data)

			// 错误集中处理，这里选择提前退出
			if err != nil {
				log.Printf("process error-%v\n", err.Error())
				return
			}
		}
	}
}

func PipeProcessBuildAndRunN(ctx context.Context, input PipeSourceFunc, maxCnt int, funcs ...PipeProcessFunc) {

	var err error = nil

	// 输入
	dataChan, err := input(ctx)
	if err != nil {
		log.Printf("source error-%v\n", err.Error())
		return
	}

	var wg = sync.WaitGroup{}
	wg.Add(maxCnt)

	// pipeline构建和执行
	// maxCnt个协程同时消费处理
	for i := 0; i < maxCnt; i++ {
		go func() {
			defer wg.Done()

			var err error = nil

			for data := range dataChan {
				// 依次执行函数，上一个函数返回结果当做下个函数参数
				for _, processFunc := range funcs {
					data, err = processFunc(ctx, data)

					// 错误集中处理，这里选择提前退出
					if err != nil {
						log.Printf("process error-%v\n", err.Error())
						return
					}
				}
			}
		}()
	}

	wg.Wait()
}
