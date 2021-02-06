package main

import (
	"container/list"
	"fmt"
)

//遍历list各个节点
func ShowList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	// 创建list
	l := list.New()

	//清空list
	l.Init()

	// 插入元素
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	e2 := l.InsertAfter(2, e1)
	e3 := l.InsertBefore(3, e4)

	ShowList(l)

	//首尾元素
	fmt.Printf("Len: %d Front Element:%d Front Element:%d\n", l.Len(), l.Front().Value, l.Back().Value)

	//演示MoveAfter,MoveBefore,MoveToBack,MoveToFront
	fmt.Println("==================after call MoveAfter==================")
	l.MoveAfter(e2, e3)
	ShowList(l)

	fmt.Println("==================after call MoveBefore==================")
	l.MoveBefore(e2, e3)
	ShowList(l)

	fmt.Println("==================after call MoveToBack==================")
	l.MoveToBack(e1)
	ShowList(l)

	fmt.Println("==================after call MoveToFront==================")
	l.MoveToFront(e1)
	ShowList(l)

	fmt.Println("==================after call Remove==================")
	l.Remove(e1)
	l.Remove(e4)
	ShowList(l) //2,3
}
