package utils

import (
	"net/http"
	"sync"
	"time"

	"../models"
	uuid "github.com/satori/go.uuid"
)

const (
	cookieName    = "go_session"
	cookieExpires = 24 * 2 * time.Hour
)

//Session con esto almacenamos un usuario que inicio session
var Session = struct {
	m map[string]*models.Profesor
	sync.RWMutex
}{m: make(map[string]*models.Profesor)}

//SetSesion crea un cookie de inicio de sesion para el usuario
func SetSesion(profe *models.Profesor, w http.ResponseWriter) {
	Session.Lock() //con esto bloque qu una go rutine intente acceder a esta variable Session
	defer Session.Unlock()
	u2, _ := uuid.NewV4()
	Session.m[u2.String()] = profe

	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   u2.String(),
		Path:    "/",
		Expires: time.Now().Add(cookieExpires),
	}
	http.SetCookie(w, cookie)
}

//DeleteSesion elimina la cookie del usuario
func DeleteSesion(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "cookie_value",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request) string {
	if cookie, err := r.Cookie(cookieName); err == nil {
		return cookie.Value
	}
	return ""

}

//EsAutentificado retorna un valor booleano si una cookie se encuentra
func EsAutentificado(r *http.Request) bool {

	return getValCookie(r) != ""
}

//GetProfesor obtiene un profesor que haya iniciado sesion mediante la cookie
func GetProfesor(r *http.Request) *models.Profesor {
	Session.Lock()
	defer Session.Unlock()

	uid := getValCookie(r)
	if profe, ok := Session.m[uid]; ok {
		return profe
	}
	return &models.Profesor{}
}
