package batchCommit

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func sendDoneSignal(done chan bool) {
	time.Sleep(30 * time.Second)
	done <- true
}

func TestCommitLog(t *testing.T) {
	var (
		logBatch        []string
		commitTimeout   = 5
		CommitBatchSize = 5
	)

	dataCh := make(chan string, 100)

	//30秒后任务完成，发送完成通知
	done := make(chan bool)
	go sendDoneSignal(done)

	//1.写协程，每隔1秒往channel中写数据
	index := 0
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	//2.发送协程
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("write done! time:", time.Now().String())
				return
			case <-ticker.C:
				index++

				if (index%CommitBatchSize == 0 && rand.Int()%2 == 1) || index%CommitBatchSize != 0 { //write data to channel
					str := strconv.Itoa(index)
					dataCh <- str
				}
			}
		}
	}()

	//2.处理协程，从channel中读取数据，然后根据两种方式进行提交commit操作
	duration := time.Duration(commitTimeout) * time.Second
	CommitTimer := time.NewTicker(duration)
	defer CommitTimer.Stop()
	go func() {
		for {
			select {
			case data, ok := <-dataCh:
				if !ok { //数据通道已关闭，没数据了
					//如果对面不关闭dataCh通道，则ok一直为true
					fmt.Println("data channel is no data!")
					CommitTimer.Stop()
					return
				}

				//2.1 将数据append到容器中
				logBatch = append(logBatch, data)

				//2.2 缓存的数量达到一定阈值
				if len(logBatch) >= CommitBatchSize {
					//批次满了，进行发送
					fmt.Println("commit by reach CommitBatchSize,logBath size:", len(logBatch))

					//清空logBatch数据
					logBatch = nil

					//让定时器重新计时
					CommitTimer.Reset(duration)
				}
			case <-CommitTimer.C: //2.2 timeout时间超时机制
				fmt.Println("commit by reach CommitTimer,logBath size:", len(logBatch), time.Now().String())

				//清空logBatch数据
				logBatch = nil
			case <-done:
				fmt.Println("receive done signal!")
				CommitTimer.Stop()
				return
			}
		}
	}()

	time.Sleep(40 * time.Second)
}
