package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func main() {
	p := Person{
		name: "test",
		age:  20,
	}
	fmt.Println(p.name, p.age)
}
