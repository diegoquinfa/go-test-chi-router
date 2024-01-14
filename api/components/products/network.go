package products

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", getProduct)
	r.Get("/{id}", getProductById)
	r.Post("/", post)

	fmt.Println("[Router] products")
	return r
}
