package main

import (
	"math/rand"
	"time"
	"log"
	"sync"
)

//the receiver says "please stop sending more" by closing an additional signal channel
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 10000
	const NumSenders = 2

	wgReceives := sync.WaitGroup{}
	wgReceives.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	//senders,one sender create a goroutine
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
					//这里必须加上default语句，否则会流程会一直阻塞在上面的case <- stopCh这
					//无法执行到下一个select语句
				default: //comment this line may cause panic
				}

				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(MaxRandomNumber):
				}
			}
		}()
	}

	//one receiver
	go func() {
		defer wgReceives.Done()
		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				//close stopChan additional channel to modify senders
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()
	wgReceives.Wait()
}
