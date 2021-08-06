package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Node struct {
	name    string
	age     uint32
	address string
}

func main() {
	//未完待续，系统的包必须要很清楚
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var node Node
	err = json.NewDecoder(resp.Body).Decode(&node)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("node:", node)
}
