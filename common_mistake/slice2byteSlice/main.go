package main

import "fmt"

func main() {
	//当把字符串转换为一个byte slice时(或反之),你就得到了一个原始数据的完整拷贝
	originStr := "text" //string
	//originStr[0] = 'T'
	newStr := []byte(originStr)
	newStr[0] = 'T'
	fmt.Printf("originStr:%s\n", originStr)
	fmt.Println("newStr:", string(newStr))
}
