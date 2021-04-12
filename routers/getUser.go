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
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
	}
	user, err := users.CheckUserID(id)
	if err != nil {
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusAccepted)
	}
}
