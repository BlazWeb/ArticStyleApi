package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Artic-Dev/ArticStyleApi-GO/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Controller() {
	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()
	router.HandleFunc("/", routers.Index)
	// Rutas gestion de usuario
	router.HandleFunc("/ban-user", routers.BanUser).Methods("POST")
	router.HandleFunc("/user", routers.RegisterUser).Methods("POST")
	router.HandleFunc("/user/{id}", routers.GetUser).Methods("GET")
	router.HandleFunc("/user/followers/{id}", routers.UserFollowers).Methods("GET")
	router.HandleFunc("/user/{id}", routers.DelUser).Methods("DELETE")
	router.HandleFunc("/auth", routers.Login).Methods("POST")
	router.HandleFunc("/update-user/{id}", routers.EditUser).Methods("PUT")

	// Rutas gestion de estilos creados
	router.HandleFunc("/create-style/{iduser}", routers.RegisterStyle).Methods("POST")
	router.HandleFunc("/save-style", routers.StyleSave).Methods("POST")
	router.HandleFunc("/style-user/{iduser}", routers.GetStylesSavedUser).Methods("GET")
	router.HandleFunc("/style/{id}", routers.GetStyle).Methods("GET")
	router.HandleFunc("/styles/{id}", routers.GetStylesUser).Methods("GET")

	// Rutas gestion de post
	router.HandleFunc("/create-post", routers.CreatePost).Methods("POST")
	router.HandleFunc("/posts-profile", routers.GetPostsProfile).Methods("GET")

	// Filtros explorador de estilos

	// Conexion API
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handlers := cors.AllowAll().Handler(router)
	log.Println("Servidor funcionando por el puerto :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))
}

func Ascii() {
	fmt.Println("")
	fmt.Println("██████████████ ████████████████ ██████████████ ██████████ ██████████████ ████████████   ██████████████ ██████  ██████")
	fmt.Println("██░░░░░░░░░░██ ██░░░░░░░░░░░░██ ██░░░░░░░░░░██ ██░░░░░░██ ██░░░░░░░░░░██ ██░░░░░░░░████ ██░░░░░░░░░░██ ██░░██  ██░░██")
	fmt.Println("██░░██████░░██ ██░░████████░░██ ██████░░██████ ████░░████ ██░░██████████ ██░░████░░░░██ ██░░██████████ ██░░██  ██░░██")
	fmt.Println("██░░██  ██░░██ ██░░██    ██░░██     ██░░██       ██░░██   ██░░██         ██░░██  ██░░██ ██░░██         ██░░██  ██░░██")
	fmt.Println("██░░██████░░██ ██░░████████░░██     ██░░██       ██░░██   ██░░██         ██░░██  ██░░██ ██░░██████████ ██░░██  ██░░██")
	fmt.Println("██░░░░░░░░░░██ ██░░░░░░░░░░░░██     ██░░██       ██░░██   ██░░██         ██░░██  ██░░██ ██░░░░░░░░░░██ ██░░██  ██░░██")
	fmt.Println("██░░██████░░██ ██░░██████░░████     ██░░██       ██░░██   ██░░██         ██░░██  ██░░██ ██░░██████████ ██░░██  ██░░██")
	fmt.Println("██░░██  ██░░██ ██░░██  ██░░██       ██░░██       ██░░██   ██░░██         ██░░██  ██░░██ ██░░██         ██░░░░██░░░░██")
	fmt.Println("██░░██  ██░░██ ██░░██  ██░░██████   ██░░██     ████░░████ ██░░██████████ ██░░████░░░░██ ██░░██████████ ████░░░░░░████")
	fmt.Println("██░░██  ██░░██ ██░░██  ██░░░░░░██   ██░░██     ██░░░░░░██ ██░░░░░░░░░░██ ██░░░░░░░░████ ██░░░░░░░░░░██   ████░░████  ")
	fmt.Println("██████  ██████ ██████  ██████████   ██████     ██████████ ██████████████ ████████████   ██████████████     ██████    ")
	fmt.Println("                                                                                                                       ")
	fmt.Println("\n                                                   Artic Solutions 2021")
	fmt.Println("")
}
