package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Artic-Dev/ArticStyleApi-GO/middlewares"
	"github.com/Artic-Dev/ArticStyleApi-GO/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Controller() {
	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()
	router.HandleFunc("/", routers.Index)
	// Rutas gestion de usuario
	router.HandleFunc("/user", routers.RegisterUser).Methods("POST")
	router.HandleFunc("/user/{id}", middlewares.ValidateJWT(routers.GetUser)).Methods("GET")
	router.HandleFunc("/user/{id}", middlewares.ValidateJWT(routers.DelUser)).Methods("DELETE")
	router.HandleFunc("/auth", routers.Login).Methods("POST")

	// Rutas gestion de estilos
	router.HandleFunc("/create-style/{iduser}", routers.RegisterStyle).Methods("POST")

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
