package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Service struct{}

func (s *Service) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func main() {
	// 实例化server
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}

	// 注册处理逻辑handler
	err = rpc.RegisterName("Service", &Service{})
	if err != nil {
		fmt.Println(err)
	}

	//启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
