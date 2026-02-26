package middleware

import "net/http"

type Wrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *Wrapper) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
