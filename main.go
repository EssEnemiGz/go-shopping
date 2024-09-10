package main

import (
	"fmt"
	"go-shopping/microservices/auth"
	"go-shopping/microservices/render"
	"net/http"
)

func index_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Golang Shop")
}

func main() {
	http.HandleFunc("/", index_page)
	http.HandleFunc("/auth", auth.Auth)
	http.HandleFunc("/render", render.Render)

	fmt.Println("Server started")
	err := http.ListenAndServe(":5001", nil)
	fmt.Println(err)
}
