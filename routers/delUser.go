package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/gorilla/mux"
)

type sendmessage struct {
	Message string `json:"message"`
	Status  bool   `json:"success"`
}

func DelUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idAsString := mux.Vars(r)["id"]
	id, err := helpers.StringToInt64(idAsString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = users.DelUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err == sql.ErrNoRows {
			var txt = "Not search user"
			e := sendmessage{txt, false}
			json.NewEncoder(w).Encode(e)
			return
		}
		e := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(e)
		panic(err)
	}
	e := sendmessage{"Deleted completed successfully", true}
	json.NewEncoder(w).Encode(e)
	w.WriteHeader(http.StatusAccepted)
}
