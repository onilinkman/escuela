package utils

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("t").ParseGlob("./templates/**/*.html"))

var errorTemplate = template.Must(template.ParseFiles("templates/error.html"))

//RenderErrorTemplate carga la pagina de error
func RenderErrorTemplate(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorTemplate.Execute(w, nil)
}

//RenderTemplate renderiza el template
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Println(err)
		RenderErrorTemplate(w, http.StatusInternalServerError)
	}
}
