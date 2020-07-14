package models

//UnidadEducativa la estructura de la unidad educativa
type UnidadEducativa struct {
	Idue      int    `json:"idue"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Correo    string `json:"correo"`
}

var queryue = `CREATE TABLE if NOT EXISTS unidadEducativa(
	idue int primary KEY not null AUTO_INCREMENT,
	nombre varchar(60) not null,
	direccion varchar(45) not null,
	correo varchar(45) not null
)`

//CrearUE crea un onbjeto de unidad educativa
func CrearUE() {
	EjecutarExec(queryue)
}

//IngresarUE ingresa elementos a unidad educativa
func (eu *UnidadEducativa) IngresarUE() {
	query := `insert into unidadEducativa (nombre,direccion,correo) values (?,?,?)`
	EjecutarExec(query, &eu.Nombre, &eu.Direccion, &eu.Correo)
}
