package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var (
		logBatch []string
		//commitTimeout   = 3
		CommitBatchSize = 10
	)

	dataCh := make(chan string, 100)

	//10秒后任务完成，发送完成通知
	done := make(chan bool)
	go func() {
		time.Sleep(30 * time.Second)
		done <- true
	}()

	//1.写协程，每隔1秒往channel中写数据
	index := 1
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("write done! time:", time.Now().String())
				return
			case <-ticker.C:
				//write data to channel
				str := strconv.Itoa(index)
				dataCh <- str
				index++
			}
		}
	}()

	//2.处理协程，从channel中读取数据，然后根据两种方式进行提交commit操作
	CommitTimer := time.NewTicker(5 * time.Second)
	defer CommitTimer.Stop()
	go func() {
		for {
			select {
			case data := <-dataCh:
				//2.1 将数据append到容器中
				logBatch = append(logBatch, data)

				//2.2 缓存的数量达到一定阈值
				if len(logBatch) >= CommitBatchSize {
					//批次满了，进行发送
					fmt.Println("commit by reach CommitBatchSize,logBath size:", len(logBatch))
				}
			case <-CommitTimer.C: //2.2 timeout时间超时机制
				fmt.Println("commit by reach CommitTimer,logBath size:", len(logBatch), time.Now().String())
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
