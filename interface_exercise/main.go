package main

import (
	"fmt"
	"go_study/interface_exercise/realRetriever"
	"go_study/interface_exercise/mockretriever"
	"time"
)

//总结：接口变量中有实现者的类型，有实现者的值
//接口变量中自带指针
//接口变量同样采用值传递，几乎不需要使用接口的指针
//接收者是指针类型

type Retriever interface {
	Get(url string) string
}

func main() {
	var r Retriever

	//1.在r的内部，是有它的类型和内容的
	r = &realRetriever.Retriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}
	inspect(r)

	r = mockretriever.Retriever{"baidu"}
	inspect(r)

	//type assertion
	//将接口类型转换为具体类型
	if mockR, ok := r.(mockretriever.Retriever); ok {
		fmt.Println("mock Retriever:", mockR.Contents)
	} else {
		fmt.Println("not a mock Retriever")
	}

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	//获取接口的类型
	switch v := r.(type) {
	case *realRetriever.Retriever: //r此时保存的是指针
		fmt.Println("useragent:", v.UserAgent)
	case mockretriever.Retriever:
		fmt.Println("contents:", v.Contents)
	}
}
