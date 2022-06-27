package user_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/Auth-Middleware-go/data"
	"github.com/ziauddinmohammad/Auth-Middleware-go/utils"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	route_vars := mux.Vars(r)
	username := route_vars["username"]

	var user_search data.UserSearch
	user_search.Username = username

	token, err := utils.ParseToken(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	payload_map, _ := utils.ParsePayloadFromToken(token)
	//check if user is requesting his profile or not
	if payload_map["username"] != username {
		fmt.Println(payload_map[username], username)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("not authorised to access " + username))
		return
	}

	//get user profile
	user_profile, exists := data.GetUser(user_search)
	if exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(user_profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error decoding user profile"))
			return
		}
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("no user data"))
		return
	}
}
