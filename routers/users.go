package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
	"github.com/gorilla/mux"
)

// Borrar Usuario mediante el ID correspiende al usuario

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
		e := sendmessage{err.Error(), false}
		if err == sql.ErrNoRows {
			var txt = "Not search user"
			e := sendmessage{txt, false}
			json.NewEncoder(w).Encode(e)
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	e := sendmessage{"Deleted completed successfully", true}
	json.NewEncoder(w).Encode(e)
	w.WriteHeader(http.StatusAccepted)
}

// Obtiene los datos de un usuario a traves de la id

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idAsString := mux.Vars(r)["id"]
	id, err := helpers.StringToInt64(idAsString)
	if err != nil {
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
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

// Ruta para registar un nuevo usuario

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	// Validación global de repuesta
	if err != nil {
		http.Error(w, "Error in request data: "+err.Error(), 400)
		return
	}

	// Campos válidos
	validSpace := regexp.MustCompile("^[A-Za-zñ ]*$")
	valid := regexp.MustCompile("^[A0-Za9-zñ]*$")
	// Validación de datos recibidos
	if len(t.Username) < 4 {
		send := sendmessage{"Minimo cuatro carácteres", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		send := sendmessage{"The field Email is required", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(t.Name) == 0 {
		send := sendmessage{"The field Name is required", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(t.LastName) == 0 {
		send := sendmessage{"The field Last Name is required", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		send := sendmessage{"The field Password is required", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Comprobación si hay existencia de carácteres inválidos
	check := valid.MatchString(t.Username)
	if !check {
		send := sendmessage{"Carácteres inválidos", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	check = validSpace.MatchString(t.Name)
	if !check {
		send := sendmessage{"Carácteres inválidos", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	check = validSpace.MatchString(t.LastName)
	if !check {
		send := sendmessage{"Carácteres inválidos", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := helpers.CheckRealEmail(t.Email)
	if err != nil {
		http.Error(w, "Context error in: "+err.Error(), 400)
		return
	}
	if res != "valid" {
		send := sendmessage{"The email is fake", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	// Validacion de usuario inexistente
	res, err = users.CheckUserName(t.Username)
	if err != nil {
		send := sendmessage{"Error check user: " + err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if res == "no_valid" {
		send := sendmessage{"User allready", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Validate Email
	res, err = users.CheckUserEmail(t.Email)
	if err != nil {
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if res == "0" {
		send := sendmessage{"Email allready", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusCreated)
		return
	}

	// Insertar en Base de Datos
	status, err := users.InsertRegisterUser(t.Username, t.Email, t.Name, t.LastName, t.Password)
	if err != nil {
		send := sendmessage{"Error: " + err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !status {
		send := sendmessage{"No se ha logrado registrar el usuario", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Finalización exitosa
	w.WriteHeader(http.StatusCreated)

}
