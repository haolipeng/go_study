package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//测试正常输出
func TestContextDeadline(t *testing.T) {
	timeout := time.Duration(time.Second * 5)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(timeout))

	go watch(ctx, "测试协程")

	//等待10秒钟
	time.Sleep(time.Second * 10)

	fmt.Println("test Context Deadline　finished")
}
