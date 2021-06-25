package selectTest

import (
	"fmt"
	"testing"
)

func TestDeadLock(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	//ch1 <- "hello world"

	select {
	case <-ch1:
		fmt.Println("ch1 case")
	case <-ch2:
		fmt.Println("ch2 case")
	default: //如果没有default分支，那么会报deadlock错误
		fmt.Println("default case")
	}

	return
}
