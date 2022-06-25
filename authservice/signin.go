package authservice

import (
	"encoding/json"
	"net/http"

	"github.com/ziauddinmohammad/Auth-Middleware-go/data"
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
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("signed in successfuully"))
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Incorrect password"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user doesn't exists"))
		return
	}

}
