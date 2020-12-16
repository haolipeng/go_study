package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	//1、创建error group对象
	g := new(errgroup.Group)

	//2、声明三个url网址
	var urls []string = []string{
		"https://www.csdn.net",
		"https://www.baidu.com",
		//"https://www.golang.org/",
	}

	for _, url := range urls {
		fmt.Println(url)
		tmpUrl := url
		//3、error group 开协程去work
		g.Go(func() error {
			resp, err := http.Get(tmpUrl)
			if err == nil {
				fmt.Printf("%s status:%s\n", tmpUrl, resp.Status)
				resp.Body.Close()
			}
			return err
		})
	}

	//4、等待协程完成，成功 or 失败
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
		fmt.Println("failed fetched all URLs.")
	}
}
