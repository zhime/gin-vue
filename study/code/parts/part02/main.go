package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./test.log")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	fmt.Println(file.Stat())
}
