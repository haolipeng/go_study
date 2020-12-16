package main

import (
	"fmt"
	"sync"
)

func main() {
	//Using goroutines on loop iterator variables
	var values = []string{
		"a",
		"b",
		"c",
	}
	var wg sync.WaitGroup
	for _, val := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(val)
		}()
	}

	wg.Wait()
	fmt.Println("the program has exited")
}
