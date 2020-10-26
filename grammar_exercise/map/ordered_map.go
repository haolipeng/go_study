package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	//1.从map中提取key到slice切片中
	var keyList []string
	for k := range barVal {
		keyList = append(keyList, k)
	}

	fmt.Println(keyList)
	sort.Strings(keyList)
	fmt.Println(keyList)

	for _, v := range keyList {
		val, ok := barVal[v]
		if ok {
			fmt.Printf("%s->%d\n", v, val)
		}
	}
}
