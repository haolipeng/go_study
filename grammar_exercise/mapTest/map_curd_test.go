package mapTest

import "fmt"

func mapCurd() {
	//空map无法使用，需要显式初始化或make
	//var nilMap  mapTest[string] string
	//nilMap["name"] = "haolipeng" //panic: assignment to entry in nil mapTest

	// 1.mapTest 创建及初始化
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	// 2.遍历操作 元素是无序的
	fmt.Println("-----------遍历--------------")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//map获取元素,可用第二个值ok判断key在map中是否存在
	fmt.Println("-----------判断值是否存在--------------")
	courseName, ok := m["course"]
	fmt.Printf("value:%s,exist:%v", courseName, ok)

	//map中key不存在时，返回null
	if couseName, ok := m["couse"]; ok {
		fmt.Println(couseName)
	} else {
		fmt.Println("key does not exist")
	}

	//删除元素
	fmt.Println("-----------删除元素--------------")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	if !ok {
		fmt.Printf("%s element is removed\n", "name")
	}

	//map作为函数参数
	modifyMap(m)
	fmt.Println("-----------调用函数modifyMap后遍历--------------")
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//map是引用类型,以函数参数形式存在，函数内部修改会影响指向的同一个map
func modifyMap(m map[string]string) {
	m["course"] = "C++"
}
