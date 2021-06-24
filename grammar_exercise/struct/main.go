package main

import (
	"fmt"
	"unsafe"
)

//验证空结构体的大小
//golang中空字节大小为0，空结构体是zerobase变量，要兼容以前的变量
type emptyStruct struct {
}

func testEmptyStruct() {
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

//Args
//golang中默认是几字节对齐的
type Args struct {
	num1 int
	num2 int32
}

type Flag struct {
	first  uint16
	second uint32
	third  uint16
}

func testMemoryAlign() {
	fmt.Printf("Args size:%d\n", unsafe.Sizeof(Args{num1: 1, num2: 2}))
	fmt.Printf("Args align size:%d\n", unsafe.Alignof(Args{}))

	fmt.Printf("Flag size:%d\n", unsafe.Sizeof(Flag{
		first:  1,
		second: 2,
	}))
	fmt.Printf("Flag align size:%d\n", unsafe.Alignof(Flag{}))
}
func main() {
	testEmptyStruct()
	testMemoryAlign()
}
