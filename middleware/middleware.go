package middleware

import (
	"log"
	"net/http"

	"github.com/ziauddinmohammad/Auth-Middleware-go/data"
	"github.com/ziauddinmohammad/Auth-Middleware-go/jwt"
	"github.com/ziauddinmohammad/Auth-Middleware-go/utils"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Get the token and check if token is in proper format or not by parsing
		var token string
		token1, err1 := utils.ParseToken(r)
		log.Println("token is ", token1)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err1.Error()))
			return
		}
		token, err := utils.ParseToken(r)
		log.Println("token is ", token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println("token is in valid format")

		// check if token is present in token base or not
		_, exists := data.IsTokenExists(token)
		if !exists {
			log.Println("token doesnt exists in token base")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not an authorized token"))
			return
		}
		log.Println("token is present in base")

		valid, err := jwt.ValidateToken(token, data.JWT_key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized to access"))
			return
		}
		log.Println("successfully authorized")
		next.ServeHTTP(w, r)
	})
}
