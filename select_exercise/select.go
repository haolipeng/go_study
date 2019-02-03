package main

import (
	"fmt"
	"time"
)

//////////////////////select定时器////////////////////////////////
func select_timer() {
	tickTimer := time.NewTicker(1 * time.Second)
	barTimer := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-tickTimer.C:
			fmt.Println("tick")
		case <-barTimer.C:
			fmt.Println("bar")
		default:
		}
	}
}

func BkdrHash(str string) uint32 {
	var seed uint32 = 10
	var hash uint32 = 0

	for i := range str {
		hash = hash*seed + uint32(i)
	}

	return hash
}

func main() {
	var uHashValue uint32 = 0
	uHashValue = BkdrHash("www.baidu.com")

	fmt.Printf("hash value is %d", uHashValue)
}

/*
func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	var wgCnt sync.WaitGroup
	wgCnt.Add(2)

	//开启协程向通道写入数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
		wgCnt.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
		wgCnt.Done()
	}()

	//等待上面协程函数调用完成，再进行下面的逻辑处理
	wgCnt.Wait()

	for i := 0; i < 2; i++ {
		select {
		case e1 := <-ch1:
			fmt.Println("1th case is selected. first=", e1)
		case e2 := <-ch2:
			fmt.Println("2th case is selected. second=", e2)
		default:
			fmt.Println("defalut case is selected")
		}
	}

	//select和timer结合
	//select_timer()
}
*/
