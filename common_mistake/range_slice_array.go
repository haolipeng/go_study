package main

import "fmt"

func main() {
	x := []string{"a", "b", "c"}

	for _, v := range x {
		fmt.Println("value:", v)
	}
}
