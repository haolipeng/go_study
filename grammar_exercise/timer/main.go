package main

import (
	"fmt"
	"time"
)

func testTimer() {
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	fmt.Println()

	// Printed after 5 seconds
	fmt.Println("Timer is inactivated")
}

func testTimerStop() {
	timer := time.NewTimer(time.Second * 5)

	go func() {
		<-timer.C
		fmt.Println("timer inactivated")
	}()

	fmt.Println("timer is being stopping!")
	ok := timer.Stop()
	time.Sleep(time.Second * 2)
	if ok {
		fmt.Println("timer is stopped")
	}
}

func main() {
	//testTimer()
	testTimerStop()
}
