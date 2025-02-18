package sync

import (
	"context"
	"log"
	"testing"
	"time"
)

// 测试单个并发度
func TestProcessorManager_SingleConcurrency(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	m := NewProcessorManager()

	//pipeline 组装
	m.AddSource(NewTimerSource(LINES))
	m.AddSink(NewConsoleSink())

	m.AddProcessor(&SplitProcessor{})
	m.AddProcessor(&CountProcessor{})
	m.AddProcessor(&SortProcessor{})

	//定时通知退出
	go func() {
		time.Sleep(15 * time.Second)
		cancel()
	}()

	err := m.Run(ctx)
	if err != nil {
		log.Printf("Run error:%s\n", err)
		return
	}
}

// 测试多个并发度
func TestProcessorManager_MultiConcurrency(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	m := NewProcessorManager()

	//pipeline 组装
	m.AddSource(NewTimerSource(LINES))
	m.AddSink(NewConsoleSink())

	m.AddProcessor(&SplitProcessor{})
	m.AddProcessor(&CountProcessor{})
	m.AddProcessor(&SortProcessor{})

	//定时通知退出
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	const routineCnt = 5
	err := m.RunN(ctx, routineCnt)
	if err != nil {
		log.Printf("Run error:%s\n", err)
		return
	}
}
