package main

import (
	"fmt"
	"time"
	"math/rand"
)

const number = 10

func generator() chan int {
	out := make(chan int)
	go func() {

		for i := 0; i < number; i++ {
			//time.Sleep(1 * time.Second)
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
		}

		close(out)

	}()
	return out
}

//deal function从channel中取元素
func deal(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func CreateWorker() chan int {
	out := make(chan int)

	go deal(1, out)

	return out
}

func main() {
	var ch1 = generator()
	//var ch2 = generator()
	var worker = CreateWorker()

	//切片保存从channel中获取的所有元素
	values := make([]int, 10)

	for {
		var activeWorker chan int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-ch1:
			values = append(values, n)
			//case n := <-ch2:
			//values = append(values,n)
		case activeWorker <- activeValue:
			if len(values) > 1 {
				values = values[1:]
				fmt.Println("values array len is ", len(values))
			}
		}
	}
}
