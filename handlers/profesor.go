package handlers

import (
	"net/http"
	"strconv"

	"../models"
	"../utils"
)

//NuevoProfesor se obtiene los datos para ingresar un profesor a la base de datos
func NuevoProfesor(w http.ResponseWriter, r *http.Request) {
	contexto := make(map[string]interface{})
	//fmt.Println(r)
	num, err := strconv.Atoi(r.FormValue("ci"))
	if r.Method == "POST" && err == nil {
		profesor := models.Profesor{
			Ci:          num,
			Nombres:     r.FormValue("nombres"),
			Paterno:     r.FormValue("paterno"),
			Materno:     r.FormValue("materno"),
			Correo:      r.FormValue("correo"),
			Contrasenia: r.FormValue("contrasenia"),
			Idue:        1,
		}
		profesor.IngresarProfesor()

	}
	utils.RenderTemplate(w, "profesor/new", contexto)
}

//MenuProfesor enlaza la direccion hacia la pagina profesor
func MenuProfesor(w http.ResponseWriter, r *http.Request) {
	contexto := make(map[string]interface{})
	utils.RenderTemplate(w, "profesor", contexto)
}
