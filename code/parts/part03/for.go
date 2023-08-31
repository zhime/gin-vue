package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	str := "hello"
	for k, v := range str {
		fmt.Printf("index: %v, value: %c\n", k, v)
	}
}
