package main

import (
	"fmt"
	"unicode/utf8"
)

//并非所有文本都是utf-8类型
func main() {
	data := "♥"
	fmt.Println(len(data))                    //prints:3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1
}
