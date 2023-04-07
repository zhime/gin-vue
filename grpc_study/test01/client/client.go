package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立连接
	client, err := rpc.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}

	//var reply = new(string)
	var reply string
	err = client.Call("Service.Hello", "test", &reply)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}
