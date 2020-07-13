package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var username = "root"
var password = ""
var host = "127.0.0.7"
var port = 8000
var dbname = "Escuela"

func Conectar() {
	coneccion, err := sql.Open("mysql", generarURL())
	if err != nil {
		panic(err)

	} else {
		db = coneccion
	}
}

//<username>:<password>@tcp(<host>:<port>)/<database>
func generarURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
}

func Cerrar() {
	db.Close()
}
