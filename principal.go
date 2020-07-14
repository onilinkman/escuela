package main

import (
	"fmt"

	"./models"
)

func main() {
	models.Conectar()
	models.CrearUE()
	ue := models.UnidadEducativa{
		Nombre:    "unidad1",
		Direccion: "sadofja Nª124",
		Correo:    "asdasf@gmail.com",
	}
	ue.IngresarUE()
	models.Cerrar()
	fmt.Println("hola")
}
