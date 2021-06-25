package selectTest

import (
	"fmt"
	"testing"
	"time"
)

func TestForSelectBreak(t *testing.T) {
	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	timer := time.NewTimer(time.Second)
	selectCnt := 0
	for {
		select {
		case <-timer.C:
			fmt.Println("timer 触发")
			timer.Reset(time.Second)
		case <-done:
			fmt.Println("done! try break,skip for loop!")
			//这个break不会跳出for循环，select-break是类似switch-case-break的东西
			break
		}

		//验证下上面break后
		selectCnt++
		fmt.Println("call select count:", selectCnt)

		time.Sleep(1 * time.Second)
	}
}
