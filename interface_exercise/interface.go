package main

import (
	"fmt"
	"strconv"
)

//人的接口
type People interface {
	SayHi()
	Sing(lyric string)
}

type Human struct {
	name string
	age  int
	home string
}

//学生
type Student struct {
	Human
	major  string //专业
	school string //学校名称
}

//雇员
type Employee struct {
	Human
	salary  float32
	company string //公司名称
}

//Human实现接口
func (h Human) SayHi() {
	fmt.Printf("I am a Human,name is %s,age is %d\n", h.name, h.age)
}

func (h Human) Sing(lyric string) {
	fmt.Printf("Human %s Sing Song,Lyric is %s\n", h.name, lyric)
}

func (h Human) String() string {
	return h.name + " " + strconv.Itoa(h.age) + " " + h.home;
}

//Employee实现SayHi接口
func (e Employee) SayHi() {
	fmt.Printf("I am a Employee,name is %s,age is %d\n", e.name, e.age)
}

func (e Employee) Sing(lyric string) {
	fmt.Printf("Employee %s Sing Song,Lyric is %s,company is %s\n", e.name, lyric, e.company)
}

//interface{} 可以存储任意类型的数据，定义一个interface类型的变量，可以存实现这个interface的任意类型的变量
func testInterfaceSave() {
	mike := Student{Human{"mike", 29, "heilongjiang"}, "shengwu", "qinghua"}
	limei := Employee{Human{"limeimei", 21, "haerbin"}, 10000, "360"}

	//定义People类型的变量
	var i People
	i = mike
	fmt.Println("This is a Student,mike")
	i.SayHi()
	i.Sing("wo shi mike")

	//i也能存储Employee类型变量
	i = limei
	fmt.Println("This is a Employee,limeimei")
	i.SayHi()
	i.Sing("wo shi li meimei")
}

//fmt输出Human变量
func printHumanInfo() {
	h := Human{"mike", 29, "heilongjiang"}
	fmt.Println("human is ", h)
}

//interface{} use
type Element interface{}
type List []Element

func interfaceUse() {
	elemList := make(List, 3)
	elemList[0] = 1
	elemList[1] = "hello world"
	elemList[2] = Human{"haolipeng", 25, "jiang_su"}

	for index, element := range elemList {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		case Human:
			fmt.Printf("list[%d] is an Human and its value is %s\n", index, value)
		}
	}
}
func main() {
	testInterfaceSave()
	printHumanInfo()
	interfaceUse()
}
