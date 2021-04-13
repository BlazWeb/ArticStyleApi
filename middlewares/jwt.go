package middlewares

import (
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/routers"
)

/*Nos permite validar el JWT que nos viene en la petición */
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
