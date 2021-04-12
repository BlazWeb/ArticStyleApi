package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	// Validación global de repuesta
	if err != nil {
		http.Error(w, "Error in request data: "+err.Error(), 400)
		return
	}

	// Validación de datos recibidos
	if len(t.Email) == 0 {
		http.Error(w, "THe email is required", 400)
		return
	}
	if len(t.Name) == 0 {
		http.Error(w, "THe name is required", 400)
		return
	}
	if len(t.LastName) == 0 {
		http.Error(w, "THe last name is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password must be at least 6", 400)
		return
	}
	// res, err := helpers.CheckRealEmail(t.Email)
	// if err != nil {
	// 	http.Error(w, "Context error in: "+err.Error(), 400)
	// 	return
	// }
	// if res != "valid" {
	// 	http.Error(w, "The email is fake", 400)
	// 	return
	// }
	// Validacion de usuario inexistente
	res, err := users.CheckUserName(t.Username)
	if err != nil {
		http.Error(w, "Error check user: "+err.Error(), 400)
		return
	}

	if res == "no_valid" {
		http.Error(w, "User already", 400)
		return
	}
	// Validate Email
	res, err = users.CheckUserEmail(t.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if res == "0" {
		http.Error(w, "Email already", 400)
		return
	}

	// Insertar en Base de Datos
	status, err := users.InsertRegisterUser(t.Username, t.Email, t.Name, t.LastName, t.Password)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado registrar el usuario", 400)
		return
	}

	// Finalización exitosa
	w.WriteHeader(http.StatusCreated)

}
