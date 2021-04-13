package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"io"
	"os"
	"strings"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/styles"
	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/gorilla/mux"
)

func RegisterStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	// Obtiene la ID del usuario mediante la URL
	idString := mux.Vars(r)["iduser"]

	// Convierte la ID en un int
	id, err := helpers.StringToInt64(idString)

	if err != nil {
		panic(err)
	}

	// Validación de usuario existente
	err = users.CheckUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		if err == sql.ErrNoRows {
			send = sendmessage{"No user found", false}
		}
		json.NewEncoder(w).Encode(send)
		return
	}
	// Saco de Form el título
	r.ParseMultipartForm(32 << 20)
	title := r.PostFormValue("title")

	// Valido el titulo
	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Title Field is required", false}
		json.NewEncoder(w).Encode(send)
		return

	}

	// Validación global de repuesta
	if err != nil {
		http.Error(w, "Error in request data: "+err.Error(), 400)
		return
	}
	// Envio los datos a la base de datos para insertarlo, si hay un err $err o si está todo ok $t para sacar el Label
	t, err := styles.InsertRegisterStyle(id, title)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	// Obtengo del form el archivo obtengo file (archivo), handler (metadatos), err (error)
	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Saco la extensión del archivo para posteriormente hacer su validez
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/styles/" + t.Label + "." + extension

	// Validación formato permitido de archivo
	if extension != "png" && extension != "jpg" && extension != "jpeg" {
		send := sendmessage{"File not allowed", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Validacion tamaño de imagen
	if handler.Size/1024 >= 500 {
		send := sendmessage{"size can be larger than 500kb", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Abro la imagen para posteriormente copiarla en la ruta
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Error open image", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Copio el archivo
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Error copy image", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Registro completado con éxito
	w.WriteHeader(http.StatusCreated)
	send := sendmessage{"Register completed successfully!", true}
	json.NewEncoder(w).Encode(send)

}
