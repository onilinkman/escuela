package handlers

import (
	"fmt"
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

//LoginProfesor logea al Profesor
func LoginProfesor(w http.ResponseWriter, r *http.Request) {
	contexto := make(map[string]interface{})
	num, err := strconv.Atoi(r.FormValue("ci"))
	if r.Method == "POST" && err == nil {
		contrasenia := r.FormValue("contrasenia")
		profesor := models.GetProfesor(num)

		if profesor.Ci == 0 {
			fmt.Println("No existe profesor")
		} else {
			fmt.Println(profesor.Ci)
			if models.Comparar(profesor.Contrasenia, contrasenia) {
				fmt.Println("Acceso consedido")
				createCookie(w)
			}
		}

	} else {
		fmt.Println(r.Method)
	}
	utils.RenderTemplate(w, "profesor", contexto)
}

//MenuProfesor enlaza la direccion hacia la pagina profesor
func MenuProfesor(w http.ResponseWriter, r *http.Request) {
	contexto := make(map[string]interface{})
	utils.RenderTemplate(w, "profesor", contexto)
}

func createCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:  "cookie_profesor",
		Value: "cookie_value",
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}
