package main

import (
	"fmt"
	"http/4-order-api/configs"
	"http/4-order-api/internal/auth"
	"http/4-order-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}
