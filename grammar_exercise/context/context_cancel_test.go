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
