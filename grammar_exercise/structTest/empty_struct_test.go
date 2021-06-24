package structTest

import (
	"fmt"
	"testing"
	"unsafe"
)

//验证空结构体的大小
//golang中空字节大小为0，空结构体是zerobase变量，要兼容以前的变量
type emptyStruct struct {
}

func TestEmptyStruct(t *testing.T) {
	a := struct {
	}{}

	b := struct {
	}{}

	c := emptyStruct{}

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	fmt.Printf("empty struct size is %d\n", unsafe.Sizeof(struct{}{}))
}
