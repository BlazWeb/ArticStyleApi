package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/jwt"

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
	user, err := users.CheckUserIDEspecial(id)
	//userF := users.CheckUserFollowers(user.Id)

	if err != nil {
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		//json.NewEncoder(w).Encode(user)
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusAccepted)
	}
}

// Ruta para registar un nuevo usuario

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	// Saco el Form
	r.ParseMultipartForm(32 << 20)
	user := r.PostFormValue("username")
	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")
	name := r.PostFormValue("name")
	lastname := r.PostFormValue("lastname")
	birthday := r.PostFormValue("birthday")
	ip := r.PostFormValue("ip")
	birth, _ := time.Parse("2006-01-02", birthday)
	t := models.User{
		Username: user,
		Email:    email,
		Password: pass,
		Name:     name,
		LastName: lastname,
		Birthday: birth,
		IP:       ip,
	}
	// Campos válidos
	validSpace := regexp.MustCompile("^[A-Za-zñ ]*$")
	valid := regexp.MustCompile("^[A0-Za9-zñ]*$")

	// Comprobación de si esta baneado
	banned, ban, err := users.CheckBanIP(t.IP)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}

	if ban {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Estás baneado de nuestro sistema por: " + banned.Reason + " expirará el: " + banned.Duration, false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Validación de datos recibidos
	if len(t.Username) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Minimo cuatro carácteres", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	age := time.Now().Year() - t.Birthday.Year()
	if age < 16 {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Solo mayores de 16", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if len(t.Email) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"The field Email is required", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if len(t.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"The field Name is required", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if len(t.LastName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"The field Last Name is required", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if len(t.Password) < 6 {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"La contraseña debe de ser mayor a 6 carácteres", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Comprobación si hay existencia de carácteres inválidos
	check := valid.MatchString(t.Username)
	if !check {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Carácteres inválidos", false}
		json.NewEncoder(w).Encode(send)

		return
	}
	check = validSpace.MatchString(t.Name)
	if !check {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Carácteres inválidos", false}
		json.NewEncoder(w).Encode(send)

		return
	}
	check = validSpace.MatchString(t.LastName)
	if !check {
		send := sendmessage{"Carácteres inválidos", false}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(send)

		return
	}

	res, err := helpers.CheckRealEmail(t.Email)
	if err != nil {
		http.Error(w, "Context error in: "+err.Error(), 400)
		return
	}
	if res != "valid" {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"The email is fake", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	// Validacion de usuario inexistente
	res, err = users.CheckUserNameEspecial(t.Username)
	if err != nil {
		send := sendmessage{"Error check user: " + err.Error(), false}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(send)

		return
	}

	if res == "no_valid" {
		send := sendmessage{"User allready", false}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(send)

		return
	}
	// Validate Email
	res, err = users.CheckUserEmail(t.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if res == "0" {
		send := sendmessage{"Email allready", false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusCreated)
		return
	}

	// Insertar en Base de Datos
	status, err := users.InsertRegisterUser(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Error: " + err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"No se ha logrado registrar el usuario", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	// Finalización exitosa
	w.WriteHeader(http.StatusCreated)
	send := sendmessage{"Se ha creado correctamente", true}
	json.NewEncoder(w).Encode(send)

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	// Saco el Form
	r.ParseMultipartForm(32 << 20)
	user := r.PostFormValue("user")
	pass := r.PostFormValue("password")
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"El campo de usuario es obligatorio", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"El campo de contraseña es obligatorio", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	var u models.User
	u, err := users.CheckUserLogin(user, pass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"El usuario y/o la contraseña es errónea", false}
		json.NewEncoder(w).Encode(send)
		return
	}

	banned, ban, err := users.CheckBanUser(u.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}

	if ban {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"Estás baneado de nuestro sistema por: " + banned.Reason + " expirará el: " + banned.Duration, false}
		json.NewEncoder(w).Encode(send)
		return
	}

	jwtKey, err := jwt.TokenJWT(u)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}

func BanUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	// Saco el Form
	r.ParseMultipartForm(32 << 20)
	user := r.PostFormValue("username")
	reason := r.PostFormValue("reason")
	duration := r.PostFormValue("duration")
	t := models.Ban{
		Reason:   reason,
		Duration: duration,
	}
	us, err, _ := users.CheckUserExists(user)
	t.User = us.Id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		if err == sql.ErrNoRows {
			send = sendmessage{"No existe el usuario", false}
		}
		json.NewEncoder(w).Encode(send)
		return
	}

	status, err := users.BanUser(t, us.LastIP)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if status {
		w.WriteHeader(http.StatusCreated)
		send := sendmessage{"Usuario: " + us.Username + " ha sido baneado con éxito hasta el: " + t.Duration, true}
		json.NewEncoder(w).Encode(send)
		return
	}
}

func UserFollowers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idAsString := mux.Vars(r)["id"]
	id, err := helpers.StringToInt64(idAsString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	err = users.CheckUserID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"No tiene seguidores", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	userF, err := users.CheckUserFollowers(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{"No tiene seguidores", false}
		json.NewEncoder(w).Encode(send)
		return
	}
	json.Marshal(userF)
	json.NewEncoder(w).Encode(userF)
	w.WriteHeader(http.StatusAccepted)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, multipart/form-data")
	idAsString := mux.Vars(r)["id"]
	id, err := helpers.StringToInt64(idAsString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	r.ParseMultipartForm(30 << 15)
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	name := r.PostFormValue("name")
	lastname := r.PostFormValue("lastname")
	birthday := r.PostFormValue("birthday")
	//file, handler, err := r.FormFile("banner")

	birth, _ := time.Parse("2006-01-02", birthday)
	user, err := users.CheckUserIDEspecial(id)
	userl := user.Username

	// Campos válidos
	validSpace := regexp.MustCompile("^[A-Za-zñ ]*$")
	valid := regexp.MustCompile("^[A0-Za9-zñ]*$")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}

	if len(username) > 0 {
		user.Username = username
		check := valid.MatchString(user.Username)
		if !check {
			w.WriteHeader(http.StatusBadRequest)
			send := sendmessage{"Carácteres inválidos", false}
			json.NewEncoder(w).Encode(send)
			return
		}
	}
	if len(email) > 0 {
		user.Email = email
		res, err := helpers.CheckRealEmail(email)
		if err != nil {
			http.Error(w, "Context error in: "+err.Error(), 400)
			return
		}
		if res != "valid" {
			w.WriteHeader(http.StatusBadRequest)
			send := sendmessage{"The email is fake", false}
			json.NewEncoder(w).Encode(send)
			return
		}
	}
	if len(name) > 0 {
		user.Name = name
		check := validSpace.MatchString(user.Name)
		if !check {
			w.WriteHeader(http.StatusBadRequest)
			send := sendmessage{"Carácteres inválidos", false}
			json.NewEncoder(w).Encode(send)
			return
		}
	}
	if len(lastname) > 0 {
		user.LastName = lastname
		check := validSpace.MatchString(user.LastName)
		if !check {
			send := sendmessage{"Carácteres inválidos", false}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(send)
			return
		}
	}
	if len(birthday) > 0 {
		user.Birthday = birth
	}
	// Comprobación si hay existencia de carácteres inválidos

	stats, err := users.PutUser(user, userl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		return
	}
	if stats {
		w.WriteHeader(http.StatusAccepted)
		send := sendmessage{"Se ha actualizado con éxito", true}
		json.NewEncoder(w).Encode(send)
	}

}
