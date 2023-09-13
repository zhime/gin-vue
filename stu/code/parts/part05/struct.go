package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func (u *User) PrintUserName() {
	fmt.Println(u.Name)
}

type Account struct {
	User
	Name string
}

func (a *Account) PrintAccountName() {
	fmt.Println(a.Name)
}

func main() {
	user := User{
		Name:     "zhang",
		Mobile:   "13838383232",
		Password: "123456",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(jsonData)
	fmt.Println(jsonStr)

	account := Account{
		User: user,
		Name: "z",
	}
	fmt.Println(account)
	fmt.Println(account.Name)
	fmt.Println(account.User.Name)
	account.PrintUserName()
	account.PrintAccountName()
}
