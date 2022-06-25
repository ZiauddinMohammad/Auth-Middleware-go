package authservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ziauddinmohammad/Auth-Middleware-go/data"
	"github.com/ziauddinmohammad/Auth-Middleware-go/jwt"
)

func Signinhandler(w http.ResponseWriter, r *http.Request) {
	var (
		received_credentials data.UserCredentials
		saved_user           data.User
	)

	//decode request body to user
	err := json.NewDecoder(r.Body).Decode(&received_credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	//check if username or password is empty
	if received_credentials.Email == "" || received_credentials.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email/Password is missing"))
		return
	}

	var search_by_email data.UserSearch
	search_by_email.Email = received_credentials.Email
	saved_user, exists := data.GetUser(search_by_email)
	if exists {
		if saved_user.Password == received_credentials.Password {

			// Generate token and save it
			header := "H256"
			payload := map[string]string{
				"username":   saved_user.Username,
				"expires at": fmt.Sprint(time.Now().Add(5 * time.Minute).Unix()),
			}
			token, err := jwt.GenerateJWTToken(header, payload, data.JWT_key)
			if err != nil {
				//return when generate token failed
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("error generating token"))
				return
			}

			// Save the token in users_tokens
			data.Addusertoken(saved_user.Username, token)
			//return when token is generated and login is success
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
			return

		} else {
			// return when password is incorrect
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Incorrect password"))
			return
		}
	} else {
		// return when user doesn't exist
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user doesn't exists"))
		return
	}

}
