package products

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

type productStruct struct {
	Id    string `json:"id"`
	Price int    `json:"price"`
	jwt.StandardClaims
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Esto es un producto mi loco"))
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	p := productStruct{
		Id:    "Producto-" + id,
		Price: 25,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToke := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	stringToken, err := accessToke.SignedString([]byte(os.Getenv("SECRET_KEY_JWT")))

	if err != nil {
		log.Fatal("Error creando el token en productbyid", err)
	}

	w.Write([]byte(fmt.Sprintf("Producto con %s", id)))

	fmt.Println(stringToken)
}

func post(w http.ResponseWriter, r *http.Request) {
	var newProduct productStruct

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("asd"))
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Diego " + newProduct.Id))
}
