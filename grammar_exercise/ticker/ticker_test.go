package ticker

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("done") //10秒后表示完成
			return
		case <-ticker.C:
			fmt.Println("做了条鱼，哈哈")
		}
	}
}
