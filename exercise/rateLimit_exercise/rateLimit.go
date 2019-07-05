package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

/*
NewLimiter函数
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error)
当Limiter不允许n个事件发生时，WaitN将阻塞等待

参考资料
https://www.kancloud.cn/liupengjie/go/967665
*/

func main() {
	useAllowN()
	//useWaitN()
}

func useAllowN() {
	//convert Duration to Limit
	r := rate.Every(1)
	limit := rate.NewLimiter(r, 10)
	for {
		if limit.AllowN(time.Now(), 8) {
			fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
		} else {
			time.Sleep(time.Second * 2)
		}
	}
}

func useWaitN() {
	r := rate.Every(2)
	limiter := rate.NewLimiter(r, 5)
	c, _ := context.WithCancel(context.TODO())
	fmt.Println(limiter.Limit(), limiter.Burst())

	for {
		limiter.WaitN(c, 3)
		fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
	}
}
