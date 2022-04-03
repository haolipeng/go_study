package channel

import (
	"fmt"
	"testing"
	"time"
)

//TestReadFromClosedChannel
//发送数据到channel，关闭channel，从channel读取两次数据
//关闭后的通道，是否还可以从里面读出数据?
func TestReadFromClosedChannel(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 18
	close(ch) //关闭后是否可从通道中读取到18这个值

	for {
		select {
		//case data := <-ch: //错误做法,读取已关闭的通道会一直读取到类型的空值
		//正确做法，采用comma语法，判断数值是否存在
		case data, ok := <-ch:
			if ok {
				fmt.Println("获取数据", data)
			} else {
				fmt.Println("通道中无数据了!")
			}

		default:
		}
		time.Sleep(time.Second)
	}
}
