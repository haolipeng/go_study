package async

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// TimerSource 生成器 输入数据依次放入输出通道
type TimerSource struct {
	Nums []int
}

func NewTimerSource(nums ...int) *TimerSource {
	return &TimerSource{Nums: nums}
}

func (t *TimerSource) Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error) <-chan int {
	defer wg.Done()

	outChannel := make(chan int, 10)

	go func() {
		defer close(outChannel)

		i := 0

		for {
			select {
			// 定时输出
			case <-time.After(1 * time.Second):
				if i >= len(t.Nums) {
					return
				}

				s := t.Nums[i]
				i = i + 1
				if s < 0 {
					errChan <- errors.New("Invalid Num")
					continue
				}

				outChannel <- s

			// 外部中断
			case <-ctx.Done():
				log.Println("Timer source received cancel signal")
				return
			}
		}
	}()

	return outChannel
}
