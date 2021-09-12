package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type GoodOrder struct {
	a int8
	b int16
	c int32
}

type BadOrder struct {
	a int8
	c int32
	b int16
}

type ArgsSingle struct {
	b uint16
}

type Args struct {
	a int8
	b uint16
	c uint64
}

type ArgsInt struct {
	a int
	c int16
	b int
}

func TestSizeof(t *testing.T) {

	l := unsafe.Sizeof(Args{})
	fmt.Println("Args length:", l)

	l = unsafe.Sizeof(ArgsInt{})
	fmt.Println("ArgsInt length:", l)
}

func TestSizeofByOrder(t *testing.T) {
	//测试结构体成员顺序对结构体内存对齐的影响
	fmt.Println(unsafe.Sizeof(GoodOrder{})) // 8
	fmt.Println(unsafe.Sizeof(BadOrder{}))  // 12
}

func TestAlignOf(t *testing.T) {
	//对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，
	//unsafe.Alignof(x) 等于其中的最大值。
	fmt.Println(unsafe.Alignof(Args{}))

	//结构体中只有单个元素，直接计算unsafe.Alignof
	fmt.Println(unsafe.Alignof(ArgsSingle{}))

	var arr []int
	fmt.Println(unsafe.Sizeof(unsafe.Alignof(arr)))
}

/////////////////////////////////////////////////////////////////////////
type EmptyStructBefore struct {
	a struct{}
	b uint32
}

type EmptyStructAfter struct {
	b uint32
	a struct{} //当struct{}作为最后一个字段时，需要填充额外的内存保证安全。
}

func TestEmptyStruct(t *testing.T) {
	fmt.Println("before", unsafe.Sizeof(EmptyStructBefore{}))
	fmt.Println("after", unsafe.Sizeof(EmptyStructAfter{}))
}
