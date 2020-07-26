package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
	v1 "./handlers/api/v1"
	"./models"
	"./utils"
	"github.com/gorilla/mux"
)

func main() {
	models.Conectar()

	mux := mux.NewRouter()

	assets := http.FileServer(http.Dir("assets"))
	statics := http.StripPrefix("/assets/", assets)
	mux.PathPrefix("/assets/").Handler(statics)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contexto := make(map[string]interface{})
		utils.RenderTemplate(w, "inicio", contexto)
	})
	mux.HandleFunc("/profesor/login", handlers.LoginProfesor).Methods("GET", "POST")
	mux.HandleFunc("/profesor/logout", handlers.LogoutProfesor).Methods("GET")
	editHandler := handlers.Autentificacion(handlers.EditProfesor)
	editHandler = handlers.MiddlewareTwo(editHandler)
	mux.Handle("/profesor/edit", editHandler).Methods("GET", "POST")

	mux.HandleFunc("/profesor/new", handlers.NuevoProfesor).Methods("GET", "POST")

	mux.HandleFunc("/api/v1/profesor/{ci:[0-9]+}", v1.GetProfesor).Methods("GET")
	mux.HandleFunc("/api/v1/profesores/", v1.GetProfesores).Methods("GET")

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

	models.Cerrar()
	fmt.Println("hola")
}
