package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

type Option func(s *Server)

func WithProtocol(proto string) Option {
	return func(s *Server) {
		s.Protocol = proto
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxConn int) Option {
	return func(s *Server) {
		s.MaxConns = maxConn
	}
}

//NewServer 先填写上必须的字段，然后是可选字段
func NewServer(addr string, port int, opts ...Option) *Server {
	//创建Server对象，并填写可选项的默认值
	s := &Server{
		Addr:     "127.0.0.1",
		Port:     8080,
		Protocol: "udp",
		Timeout:  time.Second * 10,
		MaxConns: 10,
	}

	//都选项列表中每项都应用
	for _, option := range opts {
		option(s)
	}

	return s
}

func main() {
	s := NewServer("xxxx", 1234, WithProtocol("tcp"), WithTimeout(time.Minute), WithMaxConns(100))
	fmt.Printf("server:%v\n", s)
}

//好处如下：
//编程的一大特点是分离变化点和不变点
