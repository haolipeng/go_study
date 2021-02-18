package main

import (
	"fmt"
	"strings"
)

func main() {
	ifStr := "3: eth0@if246: <BROADCAST> mtu 1500 qdisc noqueue state up"
	//查找某个子串的索引
	index := strings.Index(ifStr, "eth0")
	fmt.Println("out for loop index =", index)

	//用":"来分割字符串
	var ifIndex string
	strArray := strings.Split(ifStr, ":")

	for _, v := range strArray {
		//fmt.Println(v)
		//是否包含子串
		if strings.Contains(v, "eth0") {
			str := strings.Split(v, "eth0@if")
			ifIndex = str[1]
			fmt.Println(ifIndex)
		}
	}
}
