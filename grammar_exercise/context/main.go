package main

import (
	"context"
	"fmt"
	"time"
)

//context的使用方法比较多
//三种场景
//1.timeout
//2.call cancel function
//3.deadline
//4.data transfer with
func testContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(4 * time.Second)
	fmt.Println("可以了，通知监控停止")

	//主动调用取消函数
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

//context设置超时时间，来控制go协程生命周期
func testContextTimeout() {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*7))
	go watch(ctx, "测试协程")

	//等待10秒钟
	time.Sleep(time.Second * 10)

	// 打印最后的时间
	oneDay := time.Now()
	dayStr := fmt.Sprintf("%04d-%02d-%02d-%02d-%02d-%02d", oneDay.Year(), oneDay.Month(), oneDay.Day(), oneDay.Hour(), oneDay.Minute(), oneDay.Second())
	fmt.Printf("最后时间 :%s\n", dayStr)
}

//测试正常输出
func testContextDeadline() {
	timeout := time.Duration(time.Second * 5)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(timeout))

	go watch(ctx, "测试协程")

	//等待10秒钟
	time.Sleep(time.Second * 10)

	fmt.Println("test Context Deadline　finished")
}

func testContextValue() {
	key := "name"
	value := "haolipeng"
	ctx := context.WithValue(context.Background(), key, value)
	go func(ctx context.Context) {
		key := "name"
		value := ctx.Value(key)
		fmt.Println(value)
	}(ctx)

	time.Sleep(time.Second * 4)
}

func main() {
	//testContextCancel()
	//testContextTimeout()
	//testContextDeadline()
	testContextValue()
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "context timeout 监控退出，停止了...")
			return
		default:
			oneDay := time.Now()
			dayStr := fmt.Sprintf("%04d-%02d-%02d-%02d-%02d-%02d", oneDay.Year(), oneDay.Month(), oneDay.Day(), oneDay.Hour(), oneDay.Minute(), oneDay.Second())
			fmt.Printf("%s goroutine监控中... time:%s\n", name, dayStr)
			time.Sleep(2 * time.Second)
		}
	}
}
