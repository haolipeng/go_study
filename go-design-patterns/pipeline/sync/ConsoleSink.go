package sync

import (
	"context"
	"errors"
	"fmt"
)

type ConsoleSink struct {
}

func NewConsoleSink() ISink {
	return &ConsoleSink{}
}

func (s *ConsoleSink) Process(ctx context.Context, params any) error {
	//检查数据类型
	if v, ok := params.([]WordCount); !ok {
		return errors.New("Sink Input type error ")
	} else {
		for i := range v {
			fmt.Printf("单词:%s - 出现次数:%d\n", v[i].word, v[i].cnt)
		}

		return nil
	}
}
