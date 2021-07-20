package main

import (
	"container/heap"
	"fmt"
)

type IpCount struct {
	address string
	count   int
}

type IpCountArray []IpCount

func (h IpCountArray) Len() int           { return len(h) }
func (h IpCountArray) Less(i, j int) bool { return h[i].count < h[j].count }
func (h IpCountArray) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IpCountArray) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(IpCount))
}

func (h *IpCountArray) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//topK功能的实现
func main() {
	h := &IpCountArray{
		{
			address: "192.168.1.1",
			count:   2,
		},
		{
			address: "192.168.1.2",
			count:   3,
		},
		{
			address: "192.168.1.3",
			count:   1,
		},
		{
			address: "192.168.1.4",
			count:   5,
		},
	}
	heap.Init(h)
	heap.Push(h, IpCount{
		address: "192.168.1.5",
		count:   4,
	})
	for h.Len() > 0 {
		fmt.Printf("%v \n", heap.Pop(h))
	}
}
