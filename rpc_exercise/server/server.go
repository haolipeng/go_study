package main

import (
	"net/rpc"
	"go_study/rpc_exercise"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		//deal with goroutine
		go jsonrpc.ServeConn(conn)
	}
}
