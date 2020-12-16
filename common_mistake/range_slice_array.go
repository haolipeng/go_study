package main

import "fmt"

func main() {
	x := []string{"a", "b", "c"}

	for _, v := range x {
		fmt.Printf("value:%v pointer:%v\n", v, &v)
	}
}
