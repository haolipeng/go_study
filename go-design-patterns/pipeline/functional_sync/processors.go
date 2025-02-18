package functional_sync

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
)

// 按行分割
func splitByLine(ctx context.Context, params any) (any, error) {
	if v, ok := params.(string); !ok {
		return nil, errors.New("Split Input type error ")
	} else {
		return strings.Split(v, "\n"), nil
	}
}

// 拆分单词并统计词频
func countByWord(ctx context.Context, params any) (any, error) {
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
func sortByCount(ctx context.Context, params any) (any, error) {
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

// 输出
func outTop3(ctx context.Context, params any) (any, error) {
	if v, ok := params.([]struct {
		cnt  int
		word string
	}); !ok {
		return nil, errors.New("Output Input type error ")
	} else {
		for i, _ := range v {
			fmt.Printf("单词:%s - 出现次数:%d\n", v[i].word, v[i].cnt)
		}

		return nil, nil
	}
}
