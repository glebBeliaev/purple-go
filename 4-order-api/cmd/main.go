package main

import (
	"http/4-order-api/configs"
	"http/4-order-api/internal/auth"
	"http/4-order-api/internal/product"
	"http/4-order-api/pkg/db"
	"http/4-order-api/pkg/logger"
	"http/4-order-api/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	database := db.NewDb(conf)
	productRepository := product.NewProductRepository(database)

	// logger
	l := logger.New()

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	handler := middleware.LoggingJSON(l)(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	l.Info("server started at :8081")

	if err := server.ListenAndServe(); err != nil {
		l.WithError(err).Fatal("server failed")
	}
}
