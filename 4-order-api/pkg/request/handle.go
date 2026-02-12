package request

import (
	"http/4-order-api/pkg/res"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err, http.StatusBadRequest)
		return nil, err
	}
	err = ValidatePayload(body)
	if err != nil {
		res.Json(*w, err, http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}
