package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Email  string
	age    int
	home   string
	school string
}

func (user User) SetName(name string) {
	user.Name = name
}

func (user User) SetAge(age int) {
	user.age = age
}

func (user User) Show() {
	fmt.Println("hello world")
}
func main() {
	u := User{
		Name:   "haolipeng",
		Email:  "1078285863@qq.com",
		age:    31,
		home:   "shuangliu",
		school: "zhongbei",
	}

	t1 := reflect.ValueOf(u)
	//获取结构体中每个成员的赋值
	fmt.Printf("method number:%d\n", t1.NumMethod())

	t1.MethodByName("Show").Call([]reflect.Value{})
}
