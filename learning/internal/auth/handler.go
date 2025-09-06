package auth

import (
	"encoding/json"
	"fmt"
	"http/learning/configs"
	"http/learning/pkg/res"
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

		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err, http.StatusBadRequest)
		}
		fmt.Println(payload)

		data := LoginPayload{
			Token: handler.Config.Auth.Token,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
