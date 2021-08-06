package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	//ReadSlice返回的[]byte是指向Reader中的buffer，而不是copy一份返回
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. It is the home 1 of gophers"))
	lines, ok := reader.ReadSlice('\n')
	if ok == bufio.ErrBufferFull || ok == io.EOF {
		fmt.Println("error buffer full!")
	}

	//lines, _ := reader.ReadBytes('\n')
	//lines, _ := reader.ReadString('\n')
	fmt.Printf("line is:%s\n", lines)

	n, _ := reader.ReadSlice('\n')
	fmt.Printf("after read option,line is:%s\n", lines)

	fmt.Println(string(n))
}
