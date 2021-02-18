package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
	content string
)

func getFileInfo(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func main() {
	//创建空文件
	fileName := "test.txt"
	newFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := newFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s file size:%d\n", fileName, fileInfo.Size())

	newFile.WriteString("hello, my name is haolipeng!")

	content, _ := ioutil.ReadFile(fileName)
	fmt.Printf("%s file content:%s\n", fileName, content)

	//2.截断文件 如截断100字节，多出的内容去掉，不足的内容用nil补全
	os.Truncate(fileName, 5)
	content, _ = ioutil.ReadFile(fileName)
	fmt.Printf("%s file content:%s\n", fileName, content)

	//3.获取文件信息
	getFileInfo(fileName)

	//4.重命名
	originPath := "test.txt"
	newPath := "test2.txt"
	err = os.Rename(originPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	//5.删除文件
	err = os.Remove(newPath)
	if err != nil {
		log.Fatal(err)
	}

	//close file
	newFile.Close()

	//快写文件，创建/打开文件，写字节slice，关闭文件操作
	sliceContent := []byte("colasoft company is best")
	ioutil.WriteFile("haolipeng.txt", sliceContent, 0666)
}
