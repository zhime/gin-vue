package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprintf(w, string(body))
	_, _ = fmt.Fprintf(w, "home")
}

func user(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "user")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "createUser")
}

func order(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "order")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/order", order)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
