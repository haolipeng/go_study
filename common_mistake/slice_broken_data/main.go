package main

import (
	"bytes"
	"fmt"
)

//https://studygolang.com/articles/14409
//完整表达式可以控制结果slice的容量
//input[low:high:max]
func main() {
	path := []byte("AAAA/BBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	//dir1 := path[:sepIndex]   //引用的底层slice是path，所以把path修改了
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	newPath := bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	fmt.Println("path =>", string(path))
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB (not ok)
	fmt.Println("new path =>", string(newPath))
}
