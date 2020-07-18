package models

import "log"

//Materia la estructura de la materia
type Materia struct {
	Idmateria int     `json:"idmateria"`
	Notamin   float32 `json:"notamin"`
	Notafinal float32 `json:"notafinal"`
	Nombre    string  `json:"nombre"`
	Gestion   string  `json:"gestion"`
}

var queryMateria = `CREATE TABLE if NOT EXISTS materia(
	idmateria int primary KEY not null,
	notamin float not null,
	notafinal float not null,
	nombre varchar(45),
	gestion varchar(25)
)`

//Materias se almacena un array de materias en esta variable
type Materias []Materia

//CrearTablaMateria crea la tabla de la estructura materia en la base de datos
func CrearTablaMateria() {
	EjecutarExec(queryMateria)
}

//IngresarMateria guarda en la base de datos una materia
func (m *Materia) IngresarMateria() {
	query := `INSERT INTO materia (idmateria,notamin,notafinal,nombre,gestion) VALUES (?,?,?,?,?)`
	EjecutarExec(query, &m.Idmateria, &m.Notamin, &m.Notafinal, &m.Nombre, &m.Gestion)
}

//GetMateria devuelve una estructura del tipo materia segun su idmateria de la base de datos
func GetMateria(idmateria int) Materia {
	query := `SELECT * FROM materia WHERE idmateria=?`
	row := EjecutarQuery(query, idmateria)
	materia := Materia{}
	if row.Next() {
		row.Scan(&materia.Idmateria, &materia.Notamin, &materia.Notafinal, &materia.Nombre, &materia.Gestion)
	} else {
		log.Println("No se encontro la materia con esa id ")
	}
	return materia
}

//GetMaterias devuelve todas las materias almacenadas en la base de datos
func GetMaterias() Materias {
	query := `SELECT * FROM materia`
	rows := EjecutarQuery(query)
	materias := Materias{}
	for rows.Next() {
		materia := Materia{}
		rows.Scan(&materia.Idmateria, &materia.Notamin, &materia.Notafinal, &materia.Nombre, &materia.Gestion)
		materias = append(materias, materia)
	}
	return materias
}

//GuardarMateria guarda los cambios de la materia en la base de datos
func (m *Materia) GuardarMateria() {
	query := `UPDATE materia SET notamin=?,notafinal=?,nombre=?,gestion=? WHERE idmateria=?`
	EjecutarExec(query, &m.Notamin, &m.Notafinal, &m.Nombre, &m.Gestion, &m.Idmateria)
}

//EliminarMateria elimina una materia de la base de datos en base a su idmateria
func EliminarMateria(idmateria int) {
	query := `DELETE FROM materia WHERE idmateria=?`
	EjecutarExec(query, idmateria)
}
