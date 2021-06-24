package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//context设置超时时间，来控制go协程生命周期
func TestContextTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*7))
	go watch(ctx, "测试协程")

	//等待10秒钟
	time.Sleep(time.Second * 10)

	// 打印最后的时间
	oneDay := time.Now()
	dayStr := fmt.Sprintf("%04d-%02d-%02d-%02d-%02d-%02d", oneDay.Year(), oneDay.Month(), oneDay.Day(), oneDay.Hour(), oneDay.Minute(), oneDay.Second())
	fmt.Printf("最后时间 :%s\n", dayStr)
}
