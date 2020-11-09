package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
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
	case *RealRetriever: //r此时保存的是指针
		fmt.Println("useragent:", v.UserAgent)
	case MockRetriever:
		fmt.Println("contents:", v.Contents)
	}
}

type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get(url string) string {
	fmt.Println("mock retriever url is ", url)

	return url
}

type RealRetriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *RealRetriever) Get(url string) string {
	resp, err := http.Get(url)
	//defer resp.Body.Close() //Unhandled error
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("response close failed")
		}
	}()

	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	return string(result)
}

func main() {
	var r Retriever

	//1.init realRetriever object
	r = &RealRetriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}

	inspect(r)

	//2.init mockRetriever object
	r = MockRetriever{"www.baidu.com"}
	inspect(r)

	//3.将接口类型转换为具体实现类型，并判断是否转换成功
	if mockR, ok := r.(MockRetriever); ok {
		fmt.Println("mock MockRetriever:", mockR.Contents)
	} else {
		fmt.Println("not a mock MockRetriever")
	}
}
