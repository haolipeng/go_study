package main

import "fmt"

func main() {
	var list []string
	list = append(list, "haolipeng")
	fmt.Println(list)

	//空的map是不允许添加元素的
	//var relation map[string]string
	//relation["platform"] = "k8s" //panic: assignment to entry in nil map
}
