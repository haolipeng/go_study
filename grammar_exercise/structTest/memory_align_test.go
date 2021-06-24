package structTest

import (
	"testing"
	"unsafe"
)

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

func TestMemoryAlign(t *testing.T) {
	t.Logf("Args size:%d\n", unsafe.Sizeof(Args{num1: 1, num2: 2}))
	t.Logf("Args align size:%d\n", unsafe.Alignof(Args{}))

	t.Logf("Flag size:%d\n", unsafe.Sizeof(Flag{
		first:  1,
		second: 2,
	}))
	t.Logf("Flag align size:%d\n", unsafe.Alignof(Flag{}))
}
