package main

import (
	"fmt"

	"./models"
)

func main() {
	models.Conectar()
	models.CrearTablaProfesor()
	/*ue := models.UnidadEducativa{
		Nombre:    "unidad1",
		Direccion: "sadofja Nª124",
		Correo:    "asdasf@gmail.com",
	}*/
	//ue.IngresarUE()

	prof := models.Profesor{
		Ci:          8337564,
		Nombres:     "Christian Scion",
		Paterno:     "Marban",
		Materno:     "Callisaya",
		Correo:      "marbanchristian@gmail.com",
		Contrasenia: "1010csmc",
		Idue:        1,
	}
	prof.IngresarProfesor()
	models.Cerrar()
	fmt.Println("hola")
}
