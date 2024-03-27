package main

import "fmt"

type User struct {
	name string
}

type Account struct {
	User
}

func main() {
	a := Account{
		//name: "account",
		User: User{
			name: "user",
		},
	}
	fmt.Println(a.name)
	fmt.Println(a.User.name)
}
