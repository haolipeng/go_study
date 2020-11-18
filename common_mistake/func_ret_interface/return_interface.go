package main

import "fmt"

//知识点
//interface变量仅在类型和值为“nil”时才为“nil”，切记
func test_interface_nil() {
	var in interface{}
	var data *byte

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)
}

//错误做法
func doit(arg int) interface{} {
	var result *struct{} = nil
	if arg > 0 {
		return &struct{}{}
	}
	return result
}

//正确做法
func doitOk(arg int) interface{} {
	if arg > 0 {
		return &struct{}{}
	} else {
		return nil
	}
}

func main() {
	test_interface_nil()
	//异常情况
	var res interface{}
	res = doit(-1)
	if res != nil {
		fmt.Println("result not excepted")
	}

	res = doitOk(-1)
	if res == nil {
		fmt.Println("result excepted")
	}
}
