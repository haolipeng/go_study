package async

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// TimerSource 定时数据生成器
// Nums 存储要依次输出的数据
type TimerSource struct {
	Nums []int
}

func NewTimerSource(nums ...int) *TimerSource {
	return &TimerSource{Nums: nums}
}

// Process 实现数据源接口
// 每隔1秒生成一个数据并发送到输出通道
// 如果数据小于0，则产生错误并跳过该数据
// 当所有数据都已发送或收到取消信号时退出
func (t *TimerSource) Process(ctx context.Context, wg *sync.WaitGroup, errChan chan error) <-chan int {
	defer wg.Done()
	outChannel := make(chan int, 10)

	go func() {
		defer close(outChannel)
		i := 0

		for i < len(t.Nums) {
			select {
			case <-time.After(1 * time.Second):
				s := t.Nums[i]
				i++
				
				// 检查数据有效性
				if s < 0 {
					errChan <- errors.New("Invalid Num")
					continue
				}

				select {
				case outChannel <- s:
					// 数据发送成功
				case <-ctx.Done():
					// 收到取消信号，但仍要确保当前数据发送出去
					log.Println("Timer source received cancel signal, sending last data")
					outChannel <- s
					return
				}

			case <-ctx.Done():
				log.Println("Timer source received cancel signal")
				return
			}
		}
		log.Println("Timer source completed normally")
	}()

	return outChannel
}
