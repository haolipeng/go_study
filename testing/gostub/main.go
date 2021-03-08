package main

import (
	"errors"
	"fmt"
	"github.com/prashantv/gostub"
	"os"
)

//场景1:给全局变量打桩
func stubValue() {
	//初始值为100
	num := 100
	fmt.Println("origin value is:", num)

	//stub打桩后的值
	stubs := gostub.Stub(&num, 150)
	fmt.Println("after stub, value is:", num)

	//还原到之前的值
	stubs.Reset()
	fmt.Println("after stub reset, value is:", num)
}

//场景2:给普通函数打桩
func stubFunc() {
	var (
		output string
	)

	//不能直接把Exec传递给gostub.StubFunc
	var Exec = Exec                                                               //very important!!!
	stubs := gostub.StubFunc(&Exec, "haolipeng", errors.New("stub error string")) //打桩

	output, _ = Exec("action") //调用函数

	fmt.Printf("output: %s\n", output) //验证结果

	stubs.Reset()
}

//场景3:给标准库函数打桩(函数返回多返回值)
func stubLibraryFunc() {
	host, err := os.Hostname()
	if err == nil {
		fmt.Printf("host:%s\n", host)
	}

	var hostName = os.Hostname
	stubs := gostub.StubFunc(&hostName, "localhost", nil)
	host, err = hostName() //import
	if err == nil {
		fmt.Printf("after stub host:%s\n", host)
	}
	defer stubs.Reset()
}

func Exec(cmd string) (string, error) {
	return "hello world", errors.New("I am error!")
}

//场景4：设置环境变量
func stubEnv() {
	stubs := gostub.New()
	stubs.SetEnv("GOSTUB_VAR", "test_value") //设置环境变量
	stubs.Reset()
}

func main() {
	//stubValue()
	//stubFunc()
	//stubLibraryFunc()
	stubEnv()
}
