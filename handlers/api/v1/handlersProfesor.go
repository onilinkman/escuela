package v1

import (
	"errors"
	"net/http"
	"strconv"

	"../../../models"
	"github.com/gorilla/mux"
)

//GetProfesor obtiene un profesor
func GetProfesor(w http.ResponseWriter, r *http.Request) {
	if profesor, err := getProfesorByCi(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, profesor)
	}
}

//GetProfesores obtiene todos los profesores
func GetProfesores(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetProfesores())
}

func getProfesorByCi(r *http.Request) (*models.Profesor, error) {
	vars := mux.Vars(r)
	num, _ := strconv.Atoi(vars["ci"])
	profesor := models.GetProfesor(num)
	if profesor.Ci == 0 {
		return nil, errors.New("la ci de profesor no se encuentra en la base de datos")
	}

	return profesor, nil
}
