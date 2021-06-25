package main

import "fmt"

//如果有效实现变量的枚举
type ByteSize float64

const (
	_           = iota //ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

//实现枚举类型对应的String()函数
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// FishType 在所需枚举值上设置go:generate 指令
//go:generate stringer -type=FishType
type FishType int

const (
	A FishType = iota
	B
	C
	D
)

func main() {
	//借助Go中的String方法的默认约定，针对定义了String()方法的类型，默认输出时候会调用该String()方法
	//var val ByteSize = 1e13
	//fmt.Println(val)

	fmt.Printf("%v,%v,%v,%v\n", KB, MB, GB, TB)

	//测试go:generate生成的枚举值
	var f1 FishType = A
	fmt.Println(f1)

	var f2 FishType = D
	fmt.Println(f2)
}
