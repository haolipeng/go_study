package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 10000
	const NumReceivers = 100

	wgReceives := sync.WaitGroup{}
	wgReceives.Add(NumReceivers)

	dataCh := make(chan int, 100)

	//the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	//receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceives.Done()

			//receive values until dataCh is closed and the value buffer queue of dataCh is empty
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}
	wgReceives.Wait()
}
