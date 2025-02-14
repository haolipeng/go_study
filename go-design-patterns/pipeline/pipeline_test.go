package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

// 莎士比亚的十四行诗
var LINES = `Shakespeare Sonnet 12
When I do count the clock that tells the time
And see the brave day sunk in hideous night
When I behold the violet past prime
And sable curls all silver'd o'er with white
When lofty trees I see barren of leaves
Which erst from heat did canopy the herd
And summer's green, all girded up in sheaves
Born on the bier with white and bristly beard
Then of thy beauty do I question make
That thou among the wastes of time must go
Since sweets and beauties do themselves forsake
And die as fast as they see others grow
And nothing 'gainst Time's scythe can make defence
Save breed, to brave him when he takes thee hence`

type ISource interface {
	Process(ctx context.Context) (<-chan any, error)
}

type ISink interface {
	Process(ctx context.Context, params any) error
}

type IProcessor interface {
	Process(ctx context.Context, params any) (any, error)
}

type ProcessorManager struct {
	source ISource
	sink   ISink
	ps     []IProcessor
}

func NewProcessorManager() *ProcessorManager {
	return &ProcessorManager{}
}

// AddProcessor add processor
func (pm *ProcessorManager) AddProcessor(processor IProcessor) {
	pm.ps = append(pm.ps, processor)
}

// AddSource add source
func (pm *ProcessorManager) AddSource(source ISource) {
	pm.source = source
}

// AddSink add sink
func (pm *ProcessorManager) AddSink(sink ISink) {
	pm.sink = sink
}

func (pm *ProcessorManager) Run(ctx context.Context) error {
	var err error

	in, err := pm.source.Process(ctx)
	if err != nil {
		return err
	}

	//pipeline processor and sink operations
	for data := range in {
		//iterate processor list
		for _, p := range pm.ps {
			data, err = p.Process(ctx, data)
			if err != nil {
				log.Printf("process err %s\n", err)
				return err
			}
		}

		//sink function
		err = pm.sink.Process(ctx, data)
		if err != nil {
			log.Printf("Sink err %s\n", err)
			return err
		}
	}
	return nil
}

func (pm *ProcessorManager) RunN(ctx context.Context, maxCnt int) error {
	in, err := pm.source.Process(ctx)
	if err != nil {
		return err
	}

	//pipeline build and run
	syncProcess := func(data any) {
		//iterate processor list
		for _, v := range pm.ps {
			data, err = v.Process(ctx, data)

			if err != nil {
				log.Printf("process err %s\n", err)
				return
			}
		}

		err := pm.sink.Process(ctx, data)
		if err != nil {
			log.Printf("sink err %s\n", err)
			return
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(maxCnt)

	// 多个协程消费同一个channel
	for i := 0; i < maxCnt; i++ {
		go func() {
			defer wg.Done()

			for data := range in {
				syncProcess(data)
			}
		}()
	}

	wg.Wait()
	return nil
}

type TimerSource struct {
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
	defer ticker.Stop() // 确保资源释放

	go func() {
		defer close(out)

		for {
			select {
			case <-ticker.C: // 使用 ticker 替代 time.After
				select {
				case out <- string(input):
					// 数据发送成功
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

type ConsoleSink struct {
}

func NewConsoleSink() ISink {
	return &ConsoleSink{}
}

func (s *ConsoleSink) Process(ctx context.Context, params any) error {
	if v, ok := params.([]struct {
		cnt  int
		word string
	}); !ok {
		return errors.New("Sink Input type error ")
	} else {
		for i := range v {
			fmt.Printf("单词:%s - 出现次数:%d\n", v[i].word, v[i].cnt)
		}

		return nil
	}
}

// 按行分割

type SplitProcessor struct{}

func (p *SplitProcessor) Process(ctx context.Context, params any) (any, error) {
	if v, ok := params.(string); !ok {
		return nil, errors.New("Split Input type error ")
	} else {
		return strings.Split(v, "\n"), nil
	}
}

// 拆分单词并统计词频
type CountProcessor struct{}

func (p *CountProcessor) Process(ctx context.Context, params any) (any, error) {
	if v, ok := params.([]string); !ok {
		return nil, errors.New("Count Input type error ")
	} else {
		wordStat := make(map[string]int)

		for _, l := range v {
			words := strings.Split(string(l), " ")
			for _, word := range words {
				if v, ok := wordStat[word]; ok {
					wordStat[word] = v + 1
				} else {
					wordStat[word] = 1
				}
			}
		}

		return wordStat, nil
	}
}

// 排序计算 倒序
type SortProcessor struct{}

func (p *SortProcessor) Process(ctx context.Context, params any) (any, error) {
	if wordStat, ok := params.(map[string]int); !ok {
		return nil, errors.New("Sort Input type error ")
	} else {
		var wordStatSlice []struct {
			cnt  int
			word string
		}

		for k, v := range wordStat {
			wordStatSlice = append(wordStatSlice, struct {
				cnt  int
				word string
			}{cnt: v, word: k})
		}

		// 倒排
		sort.Slice(wordStatSlice, func(i, j int) bool {
			return wordStatSlice[i].cnt > wordStatSlice[j].cnt
		})

		// 取TOP 3
		return wordStatSlice[:3], nil
	}
}

func TestSimpleWordCount4(t *testing.T) {
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
