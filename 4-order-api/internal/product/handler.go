package product

import (
	"http/4-order-api/pkg/request"
	"http/4-order-api/pkg/res"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductRepository *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}

	router.HandleFunc("POST /product/add", handler.Create())
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /product/", handler.GetAll())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}

		product := NewProduct(body.Name, body.Description, body.Images, body.Price)

		createdProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, createdProduct, http.StatusCreated)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1) id из URL
		idStr := r.PathValue("id")
		id64, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		id := uint(id64)

		// 2) body
		body, err := request.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}

		// 3) patch-объект
		patch := &Product{
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
			Price:       body.Price,
		}

		updated, err := handler.ProductRepository.Update(id, patch)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, updated, http.StatusOK)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id64, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		if err := handler.ProductRepository.Delete(uint(id64)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// обычно на delete возвращают 204 без тела
		w.WriteHeader(http.StatusNoContent)
	}
}

func (handler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := handler.ProductRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, products, http.StatusOK)
	}
}
