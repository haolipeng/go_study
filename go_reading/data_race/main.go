package main

import (
	"runtime"
	"sync/atomic"
)

type foo struct {
	bar int64
}

func (f *foo) doFork(depth int) {
	atomic.StoreInt64(&f.bar, 1)
	defer atomic.StoreInt64(&f.bar, 0)
	if depth > 0 {
		for i := 0; i < 2; i++ {
			f2 := &foo{}
			go f2.doFork(depth - 1)
		}
	}
	runtime.GC()
}

func main() {
	f := &foo{}
	f.doFork(11)
}
