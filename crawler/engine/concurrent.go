package engine

import (
	"go_study/crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler
	WorkCount int
	ItemChan  chan Item
}

//优化接口,每个worker有一个channel，和多个worker共用一个channel，
// 就是simpleScheduler和concurrentScheduler的区别
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	GetWorkerChan() chan Request //我有一个worker，给我哪个channel
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
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
		e.createWorker(e.Scheduler.GetWorkerChan(), outCh, e.Scheduler)
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
		//not print here,print in itemSaver module
		//fmt.Printf("Got Item #%d: %v\n", itemCount, item)
		//itemCount++

		//保存item
		go func() {
			e.ItemChan <- item
		}()
	}

	//将解析过的url重新投递到管道中
	for _, request := range result.Requests {
		e.Scheduler.Submit(request)
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	//create a channel each worker
	//in := make(chan Request)//这块不自己创建channel了

	go func() {
		for {
			//tell scheduler I'm ready,nofity 类比反注册函数
			notifier.WorkerReady(in)

			//输入->并发执行处理逻辑->输出解析结果
			request := <-in

			result, err := Worker(request)
			if err != nil {
				continue
			}

			//将解析结果加入out管道
			out <- result
		}
	}()
}

func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.FetchWithUserAgent(r.Url)

	if err != nil {
		log.Printf("Fetcher:error "+
			"fetching url:%s %v", r.Url, err)
		return ParseResult{}, err
	}

	parserResult := r.ParserFunc(body)

	return parserResult, nil
}