package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("pionter:%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}

	wg.Wait()
}
