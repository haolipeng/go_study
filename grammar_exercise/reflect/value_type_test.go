package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
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

func TestTypeOf(t *testing.T) {
	var i int64
	v := reflect.TypeOf(i)
	elemType := v.Elem()
	n := elemType.NumMethod()
	for j := 0; j < n; j++ {
		m := elemType.Method(j)
		fmt.Println("method:", m.Type, m.Name)
	}
}
