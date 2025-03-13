package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	dataChan := make(chan int)

	taskNum := 3

	wg := sync.WaitGroup{}
	wg.Add(taskNum)

	// 起多个协程，data关闭时退出
	for i := 0; i < taskNum; i++ {
		// 开启go协程，持续从dataChan中读取数据
		go func(taskNo int) {
			defer wg.Done()
			fmt.Printf("Task %d run\n", taskNo)

			for {
				select {
				case _, ok := <-dataChan:
					if !ok {
						fmt.Printf("Task %d: data channel closed\n", taskNo)
						return
					}
				}
			}
		}(i)
	}

	// 通知退出
	go func() {
		time.Sleep(3 * time.Second)
		close(dataChan)
	}()

	// 等待退出完成
	wg.Wait()
}
