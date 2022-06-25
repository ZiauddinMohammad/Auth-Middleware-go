package authservice

import (
	"encoding/json"
	"net/http"

	"github.com/ziauddinmohammad/Auth-Middleware-go/data"
)

func Signuphandler(w http.ResponseWriter, r *http.Request) {
	newuser := data.User{}

	// Check if request body is valid
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	//Check if body parameters are missing
	if newuser.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email missing in request body"))
		return
	}
	if newuser.Fullname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Fullname missing in request body"))
		return
	}
	if newuser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password missing in request body"))
		return
	}
	if newuser.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username missing in request body"))
		return
	}

	if data.AddUser(newuser.Fullname, newuser.Email, newuser.Username, newuser.Password, newuser.Role) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user added successfully"))
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Email already exists"))
		return
	}
}
