package main

import "fmt"

func main() {
	//map create
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	//m2 := make(map[string] string)

	//map 遍历 hashmap 不是有序map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//map获取元素,可用第二个值ok判断key在map中是否存在
	fmt.Println("Getting Vaules")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	//map中key不存在时，返回null
	if couseName, ok := m["couse"]; ok {
		fmt.Println(couseName)
	} else {
		fmt.Println("key does not exist")
	}

	//删除元素
	fmt.Println("Deleting Values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")

	name, ok = m["name"] //declare before
	fmt.Println(name, ok)

	//map的key和value的类型

}
