package main

import (
	"fmt"

	"./models"
)

func main() {
	models.Conectar()

	models.CrearTablaDicta()

	models.Cerrar()
	fmt.Println("hola")
}
