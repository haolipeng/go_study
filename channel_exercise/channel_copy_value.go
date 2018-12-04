package main

import "fmt"

type People struct {
	name    string
	age     uint8
	Address Addr
}

type Addr struct {
	city string
}

/////////////////////////////////验证将数值投递如通道是浅拷贝还是深拷贝///////////////////////////////
func channelCopyValue() {
	p1 := People{"zhangsan", 26, Addr{"habin"}}
	var personChan = make(chan People, 1)
	fmt.Printf("p1(1):%v\n", p1)

	//将值压入channel
	personChan <- p1

	//外部修改p1的值并打印
	p1.Address.city = "chengdu"
	fmt.Printf("p2(2):%v\n", p1)

	//查看channel中的值是否被改变
	p1_copy := <-personChan
	fmt.Printf("p1_copy:%v\n", p1_copy)

	//结论：数值投入到channel中时，做了copy操作
}

func main() {
	channelCopyValue()
	fmt.Println("channel copy value lab")
}
