package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var username = "root"
var password = ""
var host = "localhost"
var port = 3306
var dbname = "Escuela"

//IngresarPassword permite ingresar un password
func IngresarPassword(passw string) {
	password = passw
}

//Conectar inicia la coneccion a la base de datos
func Conectar() {
	coneccion, err := sql.Open("mysql", generarURL())
	if err != nil {
		fmt.Println("no se pudo conectar a la base de datos")
		panic(err)

	} else {
		db = coneccion
		fmt.Println("Conexion a la base de datos exitosa!!!")
	}
}

//EjecutarQuery Ejecuta un query y devuelve filas de la tabla
func EjecutarQuery(query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

//EjecutarExec ejecuta una query y no devuelve nada
func EjecutarExec(query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
}

func generarURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
}

//Cerrar Cierra la conexion con la base de datos
func Cerrar() {
	db.Close()
}
