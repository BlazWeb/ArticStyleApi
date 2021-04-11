package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idAsString := mux.Vars(r)["id"]
	id, err := helpers.StringToInt64(idAsString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user, err := users.CheckUserID(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
