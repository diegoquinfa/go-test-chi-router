package network

import (
	"github.com/diegoquinfa/go-project/api/components/authorization"
	"github.com/diegoquinfa/go-project/api/components/products"
	"github.com/go-chi/chi/v5"
)

func AppRouter(r *chi.Mux) {
	router := chi.NewRouter()

	router.Mount("/products", products.Router())
	router.Mount("/auth", authorization.Router())

	r.Mount("/api/v1", router)
}
