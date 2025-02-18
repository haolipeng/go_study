package sync

import (
	"context"
	"errors"
	"sort"
	"strings"
)

// SplitProcessor 按行分割
type SplitProcessor struct{}

func (p *SplitProcessor) Process(ctx context.Context, params any) (any, error) {
	//校验是否为字符串类型
	if v, ok := params.(string); !ok {
		return nil, errors.New("Split Input type error ")
	} else {
		return strings.Split(v, "\n"), nil
	}
}

// CountProcessor 拆分单词并统计词频
type CountProcessor struct{}

func (p *CountProcessor) Process(ctx context.Context, params any) (any, error) {
	if v, ok := params.([]string); !ok {
		return nil, errors.New("Count Input type error ")
	} else {
		wordStat := make(map[string]int)

		//遍历切片中元素，使用空格为分隔符来拆分单词，并统计每个单词的频次
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

// SortProcessor 排序计算 倒序
type SortProcessor struct{}

// Process 处理单词统计结果，返回出现频率最高的3个单词
func (p *SortProcessor) Process(ctx context.Context, params any) (any, error) {
	// 类型断言确保输入是 map[string]int
	wordStat, ok := params.(map[string]int)
	if !ok {
		return nil, errors.New("sort Input type error")
	}

	// 预分配切片避免动态扩容带来的性能开销
	wordStatSlice := make([]WordCount, 0, len(wordStat))

	// 将map转换为切片以便排序
	for k, v := range wordStat {
		wordStatSlice = append(wordStatSlice, WordCount{cnt: v, word: k})
	}

	// 按照单词出现次数降序排序
	sort.Slice(wordStatSlice, func(i, j int) bool {
		return wordStatSlice[i].cnt > wordStatSlice[j].cnt
	})

	// 安全返回TOP 3结果
	// 如果结果少于3个，返回全部结果
	if len(wordStatSlice) > 3 {
		return wordStatSlice[:3], nil
	}
	return wordStatSlice, nil
}
