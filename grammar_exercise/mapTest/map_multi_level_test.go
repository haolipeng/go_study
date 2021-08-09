package mapTest

import (
	"fmt"
	"testing"
)

type InnerMap map[string]*Record

type Record struct {
	male bool
	name string
	age  uint32
	desc string
}

func TestMapAssign(t *testing.T) {
	var value uint32
	var ok bool
	relation := make(map[string]uint32)
	name := "haolipeng"
	relation[name] = 1
	if value, ok = relation[name]; ok {
		value = 2
	}
	fmt.Printf("value:%d\n", value)
	fmt.Printf("relation:%v\n", relation)
}

func TestMultiLevelMap(t *testing.T) {
	name := "haolipeng"
	Relations := make(map[string]InnerMap) //创建一级map

	tmpInnerMap := make(InnerMap) //创建二级map
	Relations[name] = tmpInnerMap

	r := Record{
		male: false,
		name: name,
		age:  uint32(100),
	}

	tmpInnerMap[name] = &r

	fmt.Println("origin Relations", Relations)
	for _, v := range Relations[name] {
		fmt.Printf("name:%s,age:%d,desc:%s\n", v.name, v.age, v.desc)
	}

	//重新给r对象的成员赋值
	fmt.Println("After modify desc")
	r.desc = "my name is haolipeng"
	for _, v := range Relations[name] {
		fmt.Printf("name:%s,age:%d,desc:%s\n", v.name, v.age, v.desc)
	}
}
