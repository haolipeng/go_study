package rcpsupport

import (
	"net"
	"net/rpc/jsonrpc"
	"net/rpc"
)

//后续可以直接调用rpc.Call函数
func NewClient(host string) (*rpc.Client, error) {
	conn, e := net.Dial("tcp", host)
	if e != nil {
		return nil, e
	}

	client := jsonrpc.NewClient(conn)

	return client, nil
}

func ServerRpc(host string, service interface{}) error {
	//向rpc注册服务
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		//deal with goroutine
		go jsonrpc.ServeConn(conn)
	}

	return nil
}
