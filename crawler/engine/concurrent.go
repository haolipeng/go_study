package engine

import (
	"fmt"
)

var itemCount int = 0

type ConcurrentEngine struct {
	Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(ch chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//create input channel and output channel
	//inCh := make(chan Request,10)
	outCh := make(chan ParseResult, 10)

	//configure Scheduler channel
	//e.Scheduler.ConfigureMasterWorkerChan(inCh)
	e.Scheduler.Run()

	//创建工作者协程，每个工作者对应一个request队列
	for i := 0; i < e.WorkCount; i++ {
		e.createWorker(outCh, e.Scheduler)
	}

	//添加request任务对requests队列中
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//从out channel管道获取结果集，打印并输出
	for ; ; {
		e.printParserResult(outCh)
	}
}

func (e *ConcurrentEngine) printParserResult(out chan ParseResult) {
	result := <-out

	for _, item := range result.Items {
		fmt.Printf("Got Item %d: %v\n", itemCount, item)
		itemCount++
	}

	//将解析过的url重新投递到管道中
	for _, request := range result.Requests {
		e.Scheduler.Submit(request)
	}
}

func (e *ConcurrentEngine) createWorker(out chan ParseResult, s Scheduler) {
	//create a channel each worker
	in := make(chan Request)

	go func() {
		for {
			//tell scheduler I'm ready
			e.Scheduler.WorkerReady(in)

			//输入->并发执行处理逻辑->输出解析结果
			request := <-in

			result, err := worker(request)
			if err != nil {
				continue
			}

			//将解析结果加入out管道
			out <- result
		}
	}()
}
