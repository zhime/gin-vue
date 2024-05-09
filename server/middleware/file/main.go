package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./test"

	// 判断文件夹是否存在，不存在则创建
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}

	//file, err := os.Create("./test/test.txt")
	file, err := os.OpenFile("./test/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("创建文件失败，err: %s\n", err)
	}

	msg := []byte("hello world\n")
	_, _ = file.Write(msg)

	b, _ := os.ReadFile("./test/test.txt")
	fmt.Println(string(b))

}
