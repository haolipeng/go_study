package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go Peek(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(10)
	fmt.Printf("%s\n", line)
	time.Sleep(5)
	//打开这句注释，相当于中间又进行了io read读操作，导致数据偏移发生变化,输出结果如下:
	//http://stu
	//ng.com.	 I
	fmt.Printf("%s\n", line)
}
