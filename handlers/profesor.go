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
			if models.Comparar(profesor.Contrasenia, contrasenia) {
				utils.SetSesion(profesor, w)
			}
		}

	}
	utils.RenderTemplate(w, "profesor", contexto)
}

//LogoutProfesor esta funcion elimina la cookie
func LogoutProfesor(w http.ResponseWriter, r *http.Request) {
	utils.DeleteSesion(w)
	http.Redirect(w, r, "/profesor/login", http.StatusSeeOther)
}

//EditProfesor envia a la pagina para editar al profesor
func EditProfesor(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	profe := utils.GetProfesor(r)
	if r.Method == "POST" {
		profe.Nombres = r.FormValue("nombres")
		profe.Paterno = r.FormValue("paterno")
		profe.Materno = r.FormValue("materno")
		profe.Correo = r.FormValue("correo")
		profe.Contrasenia = r.FormValue("contrasenia")
		profe.GuardarProfesor()
		return
	}
	context["nombres"] = profe.Nombres
	context["paterno"] = profe.Paterno
	context["materno"] = profe.Materno
	context["correo"] = profe.Correo
	utils.RenderTemplate(w, "profesor/edit", context)
}

//MenuProfesor enlaza la direccion hacia la pagina profesor
func MenuProfesor(w http.ResponseWriter, r *http.Request) {
	contexto := make(map[string]interface{})
	utils.RenderTemplate(w, "profesor", contexto)
}

//InicioProfesor nos direcciona a la pagina de inicio de los profesores publica
func InicioProfesor(w http.ResponseWriter, r *http.Request) {
	contexto := make(map[string]struct{})
	utils.RenderTemplate(w, "profesor/index", contexto)
}
