package models

//UnidadEducativa la estructura de la unidad educativa
type UnidadEducativa struct {
	Idue      int    `json:"idue"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Correo    string `json:"correo"`
}

//Ues Array de unidadEducava
type Ues []UnidadEducativa

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

//GetUnidadesEducativas devuelve todas las unidades educativas
func GetUnidadesEducativas() Ues {
	query := `select * from unidadEducativa`
	ues := Ues{}
	rows := EjecutarQuery(query)
	for rows.Next() {
		ue := UnidadEducativa{}
		rows.Scan(&ue.Idue, &ue.Nombre, &ue.Direccion, &ue.Correo)
		ues = append(ues, ue)
	}
	return ues
}

//GetUnidadEducativa Devuelve un objeto de unidad educativa segun su id
func GetUnidadEducativa(idue int) *UnidadEducativa {
	query := `SELECT * FROM unidadEducativa WHERE idue=?`
	rows := EjecutarQuery(query, idue)
	ue := &UnidadEducativa{}
	if rows.Next() {
		rows.Scan(&ue.Idue, &ue.Nombre, &ue.Direccion, &ue.Correo)
	}
	return ue
}

//GuardarUnidadEducativa guarda los cambios de la unidad educativa
func (eu *UnidadEducativa) GuardarUnidadEducativa() {
	query := `UPDATE unidadEducativa SET nombre=?,direccion=?,correo=? WHERE idue=?`
	EjecutarExec(query, &eu.Nombre, &eu.Direccion, &eu.Correo, &eu.Idue)
}

//EliminarUnidadEducativa elimina por id una unidad educativa
func EliminarUnidadEducativa(idue int) {
	query := `delete from unidadEducativa where idue=?`
	EjecutarExec(query, idue)
}
