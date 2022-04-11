package bufio

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	filePath = "bufioFile.txt"
)

/*
bufio.Reader实现了io.Reader接口
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

func TestBufioReader(t *testing.T) {
	//1.打开文件，只读模式
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//2.创建buffered reader,其默认缓冲区是4096字节
	bufferedReader := bufio.NewReader(file)

	// 得到字节，当前指针不变
	buffer := make([]byte, 5)
	buffer, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", buffer)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, buffer)

	//Discard 4个字节
	_, err = bufferedReader.Discard(4)
	if err != nil {
		log.Fatal(err)
	}

	// 读取一个字节, 如果读取不成功会返回Error
	oneByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", oneByte)

	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)
	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}
