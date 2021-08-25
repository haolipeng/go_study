package channelTest

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 演示如何关闭一个channel通道，以及comma语法
func TestCloseChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)

	//comma语法
	x, ok := <-ch
	fmt.Println(x, ok)
}

var counter int64

func doTask(taskChan chan string, wg *sync.WaitGroup) {
	for task := range taskChan {
		atomic.AddInt64(&counter, 1)
		fmt.Println("doTask receive task:", task)
		wg.Done()
	}
}

func TestChannelMultiReader(t *testing.T) {
	taskCount := 100
	taskChannel := make(chan string, taskCount)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go doTask(taskChannel, &wg)
	}

	for i := 0; i < taskCount; i++ {
		task := fmt.Sprintf("task-%d", i)
		taskChannel <- task
		wg.Add(1)
	}

	wg.Wait()

	//output the counter is equal to taskCount or not
	fmt.Println("counter:", counter)
}
