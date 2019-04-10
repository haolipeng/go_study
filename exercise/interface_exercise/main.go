package main

import (
	"fmt"
	"go_study/exercise/interface_exercise/mockretriever"
	"go_study/exercise/interface_exercise/realRetriever"
	"time"
)

//总结：接口变量中有实现者的类型，有实现者的值
//接口变量中自带指针
//接口变量同样采用值传递，几乎不需要使用接口的指针
//接收者是指针类型

type Retriever interface {
	Get(url string) string
}

func inspect(r Retriever) {
	//获取接口的类型
	switch v := r.(type) {
	case *realRetriever.Retriever: //r此时保存的是指针
		fmt.Println("useragent:", v.UserAgent)
	case mockretriever.Retriever:
		fmt.Println("contents:", v.Contents)
	}
}

func main() {
	var r Retriever

	//1.init realRetriever object
	r = &realRetriever.Retriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}

	inspect(r)

	//2.init mockretriever object
	r = mockretriever.Retriever{"www.baidu.com"}
	inspect(r)

	//3.将接口类型转换为具体实现类型，并判断是否转换成功
	if mockR, ok := r.(mockretriever.Retriever); ok {
		fmt.Println("mock Retriever:", mockR.Contents)
	} else {
		fmt.Println("not a mock Retriever")
	}

}
