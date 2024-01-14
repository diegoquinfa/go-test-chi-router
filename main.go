package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/diegoquinfa/go-project/api/network"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Crea el dot env perro")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	network.AppRouter(r)
	fmt.Println(os.Getenv("SECRET_KEY_JWT"))
	http.ListenAndServe(":3000", r)
}
