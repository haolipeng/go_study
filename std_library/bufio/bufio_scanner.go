package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("scanner.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	scanner.Split(bufio.ScanLines)

	// scan下一个token.
	for {
		success := scanner.Scan()
		if success == false {
			// 出现错误或者EOF是返回Error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
			} else {
				log.Fatal(err)
			}
			//出错就退出循环
			break
		}
		// 得到数据，Bytes() 或者 Text()
		fmt.Println("First word found:", scanner.Text())

		// 再次调用scanner.Scan()发现下一个token
	}
}
