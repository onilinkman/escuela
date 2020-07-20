package main

import (
	"fmt"

	"./models"
)

func main() {
	models.Conectar()

	models.CrearTablaAlumno()

	alumno := models.GetAlumno(1)
	alumno.Nombres = "christian"
	alumno.GuardarAlumno()
	fmt.Println(alumno)

	models.Cerrar()
	fmt.Println("hola")
}
