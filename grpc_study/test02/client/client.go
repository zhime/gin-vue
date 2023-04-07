package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}

	//var reply = new(string)
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("Service.Hello", "test", &reply)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}
