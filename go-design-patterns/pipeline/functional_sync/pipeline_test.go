package functional_sync

import (
	"context"
	"testing"
	"time"
)

func TestSimpleWordCount2(t *testing.T) {
	ctx := context.Background()
	PipeProcessBuildAndRun(ctx, dataSource, splitByLine, countByWord, sortByCount, outTop3)
}

func TestSimpleWordCount3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// 定时通知退出
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	PipeProcessBuildAndRunN(ctx, dataTimerSource, 3, splitByLine, countByWord, sortByCount, outTop3)
}
