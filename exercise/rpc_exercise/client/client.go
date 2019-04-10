package main

import (
	"fmt"
	"go_study/exercise/rpc_exercise"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, e := net.Dial("tcp", ":1234")
	if e != nil {
		panic(e)
	}

	client := jsonrpc.NewClient(conn)
	var result float64

	err := client.Call("DemoService.Div", rpcdemo.Args{3, 4}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
