package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAfter(t *testing.T) {
	duration := time.Minute
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("coding one second")
		case <-time.After(duration):
			fmt.Println("one minute is reach")
		}
	}
}
