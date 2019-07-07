package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func waitForTask() {
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			fmt.Println("employee:working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
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
