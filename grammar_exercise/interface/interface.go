package main

import (
	"fmt"
	"strconv"
)

type People interface {
	SayHi()
	Sing(lyric string)
}

//人类
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

//Human方法
func (h Human) SayHi() {
	fmt.Printf("I am a Human,name is %s,age is %d\n", h.name, h.age)
}

func (h Human) Sing(lyric string) {
	fmt.Printf("Human %s Sing Song,Lyric is %s\n", h.name, lyric)
}

func (h Human) String() string {
	return h.name + " " + strconv.Itoa(h.age) + " " + h.home
}

//Student并没有实现SayHi()和Sing()接口

//Employee方法
func (e Employee) SayHi() {
	fmt.Printf("I am a Employee,name is %s,age is %d\n", e.name, e.age)
}

func (e Employee) Sing(lyric string) {
	fmt.Printf("Employee %s Sing Song,Lyric is %s,company is %s\n", e.name, lyric, e.company)
}

//interface{} 可以存储任意类型的数据，
func testInterfaceSave() {
	//创建学生变量和雇员变量
	mike := Student{Human{"mike", 29, "heilongjiang"}, "shengwu", "qinghua"}
	limei := Employee{Human{"limeimei", 21, "haerbin"}, 10000, "360"}

	//定义interface类型变量（即People类型）
	var i People

	//i存储Student类型的变量
	i = mike
	fmt.Println("This is a Student,mike")
	i.SayHi()
	i.Sing("wo shi mike")

	fmt.Println("----------------------------------------------")

	//i也能存储Employee类型变量
	i = limei
	fmt.Println("This is a Employee,limeimei")
	i.SayHi()
	i.Sing("wo shi li meimei")
}

//interface{} + make 来创建内建容器
type Element interface{}
type List []Element

//List是切片类型，切片中的元素类型是interface{}
func interfaceUse() {
	elemList := make(List, 3)
	elemList[0] = 1
	elemList[1] = "hello world"
	elemList[2] = Human{"haolipeng", 25, "jiang_su"}

	for index, element := range elemList {
		//obtain interface type
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

type strategy struct {
}

var relation map[string]string = make(map[string]string)

func addStrategy() {

}

func delStrategy() {

}

func main() {
	fmt.Println("------------------testInterfaceSave---------------------------")
	testInterfaceSave()

	//fmt.Println("------------------interfaceUse---------------------------")
	//interfaceUse()
}
