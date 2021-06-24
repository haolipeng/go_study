package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//TestContextCancel
func TestContextCancel(t *testing.T) {
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
