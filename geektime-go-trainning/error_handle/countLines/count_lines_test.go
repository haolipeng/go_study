package countLines

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	//以只读形式打开文件 几种方式
	file, err := os.OpenFile("countlines.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	lines, err := CountLines(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("count lines:%d\n", lines)
}

func TestCountLinesGood(t *testing.T) {
	//以只读形式打开文件 几种方式
	file, err := os.OpenFile("countlines.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	lines, err := CountLinesGood(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("count lines:%d\n", lines)
}
