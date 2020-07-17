package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Encriptar Encripta un text y devuelve 60 caracteres
func Encriptar(text string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("no se pudo encriptar el text")
		return text
	}
	return string(hash)
}

//Comparar Compara el texto enciptado con un texto cualquiera (no encriptado) encript=texto encriptado text=texto a comparar
func Comparar(encript, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encript), []byte(text))
	return err == nil
}
