package models

//Dicta la estructura de dicta
type Dicta struct {
	Ciprofesor int `json:"ci"`
	Idrecurso  int `json:"idrecurso"`
	Idmateria  int `json:"idmateria"`
}

var queryDicta = `CREATE TABLE if NOT EXISTS dicta(
	ci_profesor int not null,
		CONSTRAINT fk_dicta_profesor
		FOREIGN KEY (ci_profesor) REFERENCES profesor(ci),
	id_recurso int not null,
		CONSTRAINT fk_dicta_recurso
		FOREIGN KEY (id_recurso) REFERENCES recurso(idrecurso), 
	id_materia int not null,
		CONSTRAINT fk_dicta_materia
		FOREIGN KEY (id_materia) REFERENCES materia(idmateria)
)`

//DictaArray un array de la estructura dicta
type DictaArray []Dicta

//CrearTablaDicta crea una tabla dicta en la base de datos
func CrearTablaDicta() {
	EjecutarExec(queryDicta)
}

//IngresarDicta ingresa a la base de datos las referencias correspondientes
func (d *Dicta) IngresarDicta() {
	query := `INSERT INTO dicta (ci_profesor,id_recurso,id_materia) VALUES (?,?,?)`
	EjecutarExec(query, &d.Ciprofesor, &d.Idrecurso, &d.Idmateria)
}

//GetRecursoDicta obtiene todos los recursos dependiendo de la id del profesor y la materia
func GetRecursoDicta(ciprof, idmateria int) Recursos {
	query := `SELECT re.* FROM recurso re,(select di.* FROM dicta di WHERE di.ci_profesor=? AND di.id_materia=?) tmp
	WHERE re.idrecurso=tmp.id_recurso`
	rows := EjecutarQuery(query, ciprof, idmateria)
	recursos := Recursos{}
	for rows.Next() {
		recurso := Recurso{}
		rows.Scan(&recurso.Idrecurso, &recurso.Titulo, &recurso.Direccion)
		recursos = append(recursos, recurso)
	}
	return recursos
}
