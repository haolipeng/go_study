package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//defer确保调用在函数结束时发生
//参数在defer语句时计算，这个是什么意思？
//defer列表先进后出，后进先出
func tryDefer() {
	//验证defer列表中元素的顺序
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(2)

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 6 {
			panic("print too many")
		}
	}
}

func writeFileError(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		err = errors.New("this is a custom error")
		// If there is an error, it will be of type *PathError.
		//判断error值的类型，分而治之
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}

		return
	}

	//close file handle
	defer file.Close()

	//buffer io
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	//write operation
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i+1)
	}

	//此时内容还在bufio缓冲区中，并没有写入磁盘中
	//writer.Flush()
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	//close file handle
	defer file.Close()

	//buffer io
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	//write operation
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i+1)
	}

	//此时内容还在bufio缓冲区中，并没有写入磁盘中
	//writer.Flush()
}

func main() {
	tryDefer()
	writeFile("test.txt")
	writeFileError("test.txt")
}
