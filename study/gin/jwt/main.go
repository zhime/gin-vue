package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func main() {
	claims := MyClaims{
		UserName: "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte("AllYourBase")
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v\n", ss, err)

	t, err := jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	fmt.Println(t.Claims.(*MyClaims).UserName)
}
