package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//多个发送者，一个接收者的例子
//可考虑增加一个传递"关闭信号"的channel，receiver通过信号channel下达关闭数据channel的指令
//senders监听到关闭信号后，停止发送数据.

//注意：go语言中，一个channel没有任何goroutine引用它，不管channel是否被关闭，都会被gc回收。
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 10

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel
	// dataCh, and its receivers are the
	// senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(index int) {
			defer func() {
				fmt.Printf("sender %d is stopped\n", index+1)
			}()
			for {
				// The try-receive operation is to try
				// to exit the goroutine as early as
				// possible. For this specified example,
				// it is not essential.
				// 通过ok来判断channel是否已关闭
				select {
				case _, ok := <-stopCh:
					if !ok {
						return
					}
				default:
				}

				// Even if stopCh is closed, the first
				// branch in the second select may be
				// still not selected for some loops if
				// the send to dataCh is also unblocked.
				// But this is acceptable for this
				// example, so the first select block
				// above can be omitted.
				select {
				case _, ok := <-stopCh:
					if !ok {
						return
					}
				case dataCh <- rand.Intn(Max):
				}
			}
		}(i)
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// The receiver of channel dataCh is
				// also the sender of stopCh. It is
				// safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()

	time.Sleep(10 * time.Second)
}
