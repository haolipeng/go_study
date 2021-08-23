package context

import (
	"context"
	"fmt"
	"time"
)

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
			//这里执行命令卡主了，即使传入context变量，也走不到case <-ctx.Done():分支
			time.Sleep(30 * time.Second)
		}
	}
}
