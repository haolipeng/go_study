package main

import (
	"fmt"
	"errors"
)

func tryRecover() {
	defer func() {
		r := recover()
		//判断recover返回值是否是error类型
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(err)
		}
	}()

	panic(errors.New("this is a error"))
}

func main() {
	tryRecover()
}
