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
	rv := []interface{}{"hi", 42, func() {}}
	for _, v := range rv {
		v := reflect.ValueOf(v)
		switch v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}
}
