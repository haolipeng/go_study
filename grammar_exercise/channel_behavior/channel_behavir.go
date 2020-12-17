package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//生产者消费者模型
func waitForTask() {
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			fmt.Println("employee:working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- fmt.Sprintf("paper %d", w)
	}

	//close channel
	close(ch)
}

func withTimeout() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
	}()

	//或者取到任务，或者是context超时控制
	select {
	case p := <-ch:
		fmt.Println("work complete", p)
	case <-ctx.Done():
		fmt.Println("moving on")
	}
}

func main() {
	waitForTask()
}
