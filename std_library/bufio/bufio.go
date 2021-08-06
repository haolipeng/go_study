package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
type Reader struct{
    buf          []byte        // 缓存
    rd           io.Reader    // 底层的io.Reader
    // r:从buf中读走的字节（偏移）；w:buf中填充内容的偏移；
    // w - r 是buf中可被读的长度（缓存数据的大小），也是Buffered()方法的返回值
    r, w         int
    err          error        // 读过程中遇到的错误
    lastByte     int        // 最后一次读到的字节（ReadByte/UnreadByte)
    lastRuneSize int        // 最后一次读到的Rune的大小(ReadRune/UnreadRune)
}
*/

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
