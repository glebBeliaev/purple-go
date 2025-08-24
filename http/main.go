package main

import (
	"fmt"
	"http/http/handler"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	handler.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
