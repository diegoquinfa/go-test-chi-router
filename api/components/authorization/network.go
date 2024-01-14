package authorization

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Post("/", validateToken)

	fmt.Println("[router] auth")
	return r
}
