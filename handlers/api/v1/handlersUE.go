package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"../../../models"
	"github.com/gorilla/mux"
)

//GetUE devuelve datos de la unidad educativa en formato json
func GetUE(w http.ResponseWriter, r *http.Request) {
	if ue, err := getUEporIdue(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, ue)
	}
}

//CrearUE crea una nueva unidad educativa mediante un json que envie un cliente
func CrearUE(w http.ResponseWriter, r *http.Request) {
	ue := &models.UnidadEducativa{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ue); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	ue.IngresarUE()
}

func getUEporIdue(r *http.Request) (*models.UnidadEducativa, error) {
	vars := mux.Vars(r)
	num, _ := strconv.Atoi(vars["idue"])
	uue := models.GetUnidadEducativa(num)
	if uue.Idue == 0 {
		return uue, errors.New("La unidad educativa no existe")
	}
	return uue, nil
}
