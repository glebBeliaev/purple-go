package auth

import (
	"fmt"
	"http/4-order-api/configs"
	"http/4-order-api/pkg/request"
	"http/4-order-api/pkg/res"
	"net/http"
)

type AuthHandler struct {
	Config *configs.Config
}
type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		fmt.Println(payload)

		data := LoginResponse{
			Token: handler.Config.Auth.Token,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		fmt.Println(payload)
	}
}
