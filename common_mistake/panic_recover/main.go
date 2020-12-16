package main

import "fmt"

//panic和recover的使用方法还需要多练习下
func main() {
	//incorrect
	/*recover()
	panic("not good")
	recover()
	fmt.Println("ok")*/

	//work
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	panic("not good")

	fmt.Println("the program is exited")
}
