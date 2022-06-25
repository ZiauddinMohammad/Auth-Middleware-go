package user_service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/Auth-Middleware-go/data"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	route_vars := mux.Vars(r)
	username := route_vars["username"]

	var user_search data.UserSearch
	user_search.Username = username
	//Check if user is authenticated or not

	//check if user is requesting his profile or not

	//get user profile
	user_profile, exists := data.GetUser(user_search)
	if exists {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(user_profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error decoding user profile"))
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("no user data"))
		return
	}
}
