package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "./handlers/api/v1"
	"./models"
	"github.com/gorilla/mux"
)

func main() {
	models.Conectar()

	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/profesor/{ci:[0-9]+}", v1.GetProfesor).Methods("GET")
	mux.HandleFunc("/api/v1/profesores/", v1.GetProfesores).Methods("GET")

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

	//models.CrearTablaRequisito()
	/*materia := models.Materia{
		Idmateria: 111,
		Notamin:   53.4,
		Notafinal: 51.0,
		Nombre:    "fundamentos de la programacion",
		Gestion:   "I/2020",
	}
	materia.IngresarMateria()

	materia2 := models.Materia{
		Idmateria: 121,
		Notamin:   60.8,
		Notafinal: 51.0,
		Nombre:    "PROGRAMACION ORIENTADA A OBJETOS",
		Gestion:   "I/2020",
	}
	materia2.IngresarMateria()*/
	/*materia2 := models.Materia{
		Idmateria: 131,
		Notamin:   51.0,
		Notafinal: 80.6,
		Nombre:    "PROGRAMACION ORIENTADA A OBJETOS",
		Gestion:   "I/2020",
	}
	materia2.IngresarMateria()

	requisito := models.Requisito{
		Idmateria:    131,
		Idmateriareq: 111,
	}
	requisito.AgregarRequisito()

	requisito2 := models.Requisito{
		Idmateria:    131,
		Idmateriareq: 121,
	}
	requisito2.AgregarRequisito()*/

	//fmt.Println(models.GetRequisitosMateria(131))

	models.Cerrar()
	fmt.Println("hola")
}
