package handlers

import (
	"net/http"

	"../utils"
)

type customHandler func(w http.ResponseWriter, r *http.Request)

//Autentificacion autentifica los datos del inicio de sesion
func Autentificacion(funcion customHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.EsAutentificado(r) {
			http.Redirect(w, r, "/profesor/login", http.StatusSeeOther)
			return
		}
		funcion(w, r)
	})
}

//MiddlewareTwo aqui se ejecutara un segundo wrap
func MiddlewareTwo(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handle.ServeHTTP(w, r)
	})
}
