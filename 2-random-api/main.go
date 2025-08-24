package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", random)
	http.ListenAndServe(":8081", nil)
}

func random(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(6) + 1
	w.Write([]byte(fmt.Sprintf("%d", num)))
}
