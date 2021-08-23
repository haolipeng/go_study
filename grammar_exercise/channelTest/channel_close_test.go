package channelTest

import (
	"fmt"
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
