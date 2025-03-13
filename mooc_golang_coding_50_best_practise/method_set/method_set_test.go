package method_set

import (
	"fmt"
	"reflect"
	"testing"
)

type Interface interface {
	M1()
	M2()
}

type T struct {
}

func (t T) M1() {
	fmt.Println("call M1() function")
}

func (t *T) M2() {
	fmt.Println("call M2() function")
}

//实现一个函数
//打印自定义类型的方法集合MethodSet
func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemType := v.Elem()
	n := elemType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!", elemType)
		return
	}
	fmt.Printf("%s's method set:\n", elemType)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemType.Method(j).Name)
	}
	fmt.Println()
}

func TestMethodSet(o *testing.T) {
	var t T
	var pt *T

	DumpMethodSet(&t) //这块为什么要传递指针
	DumpMethodSet(&pt)
	DumpMethodSet((*Interface)(nil))
}
