package sync

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

type TimerSource struct {
	//数据源
	data string
}

func (s *TimerSource) Process(ctx context.Context) (<-chan any, error) {
	// 使用带缓冲的channel避免阻塞
	out := make(chan any, 10)

	// 预先读取数据，避免重复创建Reader
	r := bufio.NewReader(strings.NewReader(s.data))
	input, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("initial read failed: %w", err)
	}

	// 创建定时器，避免重复创建
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		// 确保资源timer释放
		defer ticker.Stop()

		//协程退出时，释放channel资源
		defer close(out)

		for {
			select {
			case <-ticker.C: // 使用 ticker 替代 time.After
				select {
				case out <- string(input): // 数据发送成功
					log.Println("timer success data input!")
				default:
					// channel已满，记录日志
					log.Println("Channel full, dropping message")
				}
			//取消Source数据源
			case <-ctx.Done():
				log.Printf("Source exiting: %v", ctx.Err())
				return
			}
		}
	}()

	return out, nil
}

func NewTimerSource(data string) ISource {
	return &TimerSource{data: data}
}
