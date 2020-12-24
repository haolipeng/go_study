package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	filePath = "bufioFile.txt"
)

func bufioReader() {
	//1.打开文件，只读模式
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//2.创建buffered reader,其默认缓冲区是4096字节
	bufferedReader := bufio.NewReader(file)

	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

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

func bufioWriter() {
	//1.打开文件，只写模式
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//2.创建buffered writer,其默认缓冲区是4096字节
	bufferedWriter := bufio.NewWriter(file)

	//3.写字节到buffer
	bytesWritten, err := bufferedWriter.Write([]byte("1234567890"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten) //10字节

	//4.写字符串到buffer
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten) //16字节

	//检查缓存中字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	fmt.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	//还有多少字节可用(未使用的缓存大小)
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)

	//写内存buffer到硬盘中
	bufferedWriter.Flush()

	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("after Reset() Available buffer: %d\n", bytesAvailable)

	unflushedBufferSize = bufferedWriter.Buffered()
	fmt.Printf("after Reset() Bytes buffered: %d\n", unflushedBufferSize)

	////////////////////////重新创建一个带size的writer////////////////////////////
	//传参size 比bufferedWriter的size小时，其可用字节数不变
	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)

	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("after NewWriterSize(8000) Available buffer: %d\n", bytesAvailable)
}

func main() {
	bufioWriter()

	bufioReader()

	//////////////////////////////////////////////////////
	//1、读取最多N个字节 file.Read
	//2、读取正好N个字节
	//2、1 file.Read()可以读取一个小文件到大的byte slice中，
	// 2、2 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	//3、io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	//4、读取全部字节 ioutil.ReadAll
	//5、快读到内存 ioutil.ReadFile

}
