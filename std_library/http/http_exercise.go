package main

import (
	"fmt"
	"net/http"
)

func main() {
	//未完待续，系统的包必须要很清楚
	resp, err := http.Get("https://www.csdn.net/")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
