package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")

type AuthSession struct {
	Phone string
	Code  int
}

var sessions = map[string]AuthSession{}

type PhoneRequest struct {
	Phone string `json:"phone"`
}

type PhoneResponse struct {
	SessionID string `json:"sessionId"`
}

type VerifyRequest struct {
	SessionID string `json:"sessionId"`
	Code      int    `json:"code"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

func main() {
	http.HandleFunc("/auth/phone", handlePhone)
	http.HandleFunc("/auth/verify", handleVerify)
	http.HandleFunc("/profile", handleProfile)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePhone(w http.ResponseWriter, r *http.Request) {
	var req PhoneRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	sessionID := generateSessionID()
	code := generateCode()

	sessions[sessionID] = AuthSession{
		Phone: req.Phone,
		Code:  code,
	}

	fmt.Println("SMS code:", code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PhoneResponse{
		SessionID: sessionID,
	})
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	session, exists := sessions[req.SessionID]
	if !exists {
		http.Error(w, "session not found", http.StatusUnauthorized)
		return
	}

	if session.Code != req.Code {
		http.Error(w, "invalid code", http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(session.Phone)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	delete(sessions, req.SessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{
		Token: token,
	})
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	phone, ok := getPhoneFromToken(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "protected profile",
		"phone":   phone,
	})
}

func getPhoneFromToken(r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", false
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", false
	}

	tokenString := parts[1]

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", false
	}

	return claims.Phone, true
}

func generateJWT(phone string) (string, error) {
	claims := Claims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func generateSessionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func generateCode() int {
	return rand.Intn(9000) + 1000
}
