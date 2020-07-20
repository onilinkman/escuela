package models

//Recurso esta estructura indica la direccion de dicho recurso
type Recurso struct {
	Idrecurso int    `json:"idrecurso"`
	Titulo    string `json:"titulo"`
	Direccion string `json:"direccion"`
}

var queryRecurso = `CREATE TABLE if NOT EXISTS recurso(
	idrecurso int primary KEY not null AUTO_INCREMENT,
	titulo varchar(45) not null,
	direccion varchar(60) not null
)`

//Recursos array de recurso
type Recursos []Recurso

//CrearTablaRecurso crea la tabla de recursos
func CrearTablaRecurso() {
	EjecutarExec(queryRecurso)
}

//IngresarRecurso ingresa datos de la estructura recurso a la base de datos
func (r *Recurso) IngresarRecurso() {
	query := `INSERT INTO recurso (titulo,direccion) VALUES (?,?)`
	EjecutarExec(query, &r.Titulo, &r.Direccion)
}

//GetRecurso obtiene una estructura recurso dependiendo al id
func GetRecurso(idrecurso int) *Recurso {
	recurso := &Recurso{}
	query := `select * from recurso where idrecurso=?`
	row := EjecutarQuery(query, idrecurso)
	if row.Next() {
		row.Scan(&recurso.Idrecurso, &recurso.Titulo, &recurso.Direccion)
	}
	return recurso
}

//GetRecursos devuelve todos los recursos
func GetRecursos() Recursos {
	recursos := Recursos{}
	query := `SELECT * FROM recurso`
	rows := EjecutarQuery(query)
	for rows.Next() {
		recurso := Recurso{}
		rows.Scan(&recurso.Idrecurso, &recurso.Titulo, &recurso.Direccion)
		recursos = append(recursos, recurso)
	}
	return recursos
}

//GuardarRecurso Guarda los cambios en la estructura de recurso
func (r *Recurso) GuardarRecurso() {
	query := `UPDATE recurso SET titulo=?,direccion=? WHERE idrecurso=?`
	EjecutarExec(query, &r.Titulo, &r.Direccion, &r.Idrecurso)
}

//EliminarRecurso elimina un recurso de la base de datos
func EliminarRecurso(idrecur int) {
	query := `DELETE FROM recurso WHERE idrecurso=?`
	EjecutarExec(query, idrecur)
}
