package authorization

import (
	"fmt"
	"net/http"
)

func validateToken(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("jwt")
	if token[0:6] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("malito ese token :v"))
		return
	}
	token = token[6:]

	fmt.Println(token)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(token))
}
