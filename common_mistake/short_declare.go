package main

import "fmt"

func main() {
	x := 1
	fmt.Println("x is ", x) //1

	func() {
		x := 2
		fmt.Println("x is ", x) //2
	}()

	fmt.Println("x is ", x) //1
}
