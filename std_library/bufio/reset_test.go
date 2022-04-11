package bufio

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestReset(t *testing.T) {
	//1.构建字符串
	reader := strings.NewReader("abcd")

	//2.创建buffered reader
	bufferReader := bufio.NewReader(reader)

	//3.创建4个字节的字节缓冲区
	buffer := make([]byte, 4)

	//4.读取数据到buffer中
	_, err := bufferReader.Read(buffer)
	if err != nil {
		return
	}
	fmt.Printf("%q\n", buffer)

	//////////////////调用Reset///////////////////////////
	reader2 := strings.NewReader("xyz")
	buffer2 := make([]byte, 4)

	bufferReader.Reset(reader2)
	_, err = bufferReader.Read(buffer2)
	if err != nil {
		return
	}
	fmt.Printf("%q\n", buffer2)
}
