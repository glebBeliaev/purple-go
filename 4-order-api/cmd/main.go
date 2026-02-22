package main

import (
	"fmt"
	"http/4-order-api/configs"
	"http/4-order-api/internal/auth"
	"http/4-order-api/internal/product"
	"http/4-order-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	productRepository := product.NewProductRepository(db)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}
