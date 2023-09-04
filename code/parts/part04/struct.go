package main

import "fmt"

type Person struct {
	name   string
	age    uint8
	gender string
}

type Student struct {
	person Person
	class  string
}

func main() {
	p := Person{
		name: "test",
		age:  20,
	}

	s := Student{
		person: Person{
			name:   "dev",
			age:    18,
			gender: "男",
		},
		class: "一班",
	}
	fmt.Println(p.name, p.age)
	fmt.Println(s.person.name, s.class)
}
