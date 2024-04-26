package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	//port := "8080"

	var service map[string]int = map[string]int{
		"oms":             7779,
		"orderTemp":       8092,
		"orderConversion": 7795,
		"order":           7781,
		"omsReport":       7782,
	}

	for key, value := range service {
		fmt.Println(key, value)
		_, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(value))
		if err != nil {
			fmt.Printf("端口 %d 未被监听\n", value)
			continue
		}
		//defer conn.Close()

	}
}
