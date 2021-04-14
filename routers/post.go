package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/posts"
	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	IdProfileString := r.URL.Query().Get("profile")
	idProfile, err := helpers.StringToInt64(IdProfileString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	IdAuthorString := r.URL.Query().Get("author")
	idAuthor, err := helpers.StringToInt64(IdAuthorString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// Comprobación existencia del usuario del perfil
	err = users.CheckUserID(idProfile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		if err == sql.ErrNoRows {
			send = sendmessage{"El perfil no existe", false}
		}
		json.NewEncoder(w).Encode(send)
		return
	}
	err = users.CheckUserID(idAuthor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		if err == sql.ErrNoRows {
			send = sendmessage{"El autor no existe", false}
		}
		json.NewEncoder(w).Encode(send)
		return
	}
	r.ParseMultipartForm(30 << 25)
	content := r.PostFormValue("content")
	model := models.Post{
		Profile: idProfile,
		Author:  idAuthor,
		Content: content,
	}
	err = posts.InsertPost(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	w.WriteHeader(http.StatusCreated)
	send := sendmessage{"Se ha creado correctamente", true}
	json.NewEncoder(w).Encode(send)
}

func GetPostsProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	IdProfileString := r.URL.Query().Get("profile")
	idProfile, err := helpers.StringToInt64(IdProfileString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	model, err := posts.GetPostsProfile(idProfile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(model)
}
