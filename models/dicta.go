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
		FOREIGN KEY (ci_profesor) REFERENCES profesor(ci)
		ON DELETE NO ACTION
		ON UPDATE NO ACTION,
	id_recurso int not null,
		CONSTRAINT fk_dicta_recurso
		FOREIGN KEY (id_recurso) REFERENCES recurso(idrecurso)
			ON DELETE NO ACTION
			ON UPDATE NO ACTION, 
	id_materia int not null,
		CONSTRAINT fk_dicta_materia
		FOREIGN KEY (id_materia) REFERENCES materia(idmateria)
			ON DELETE NO ACTION
			ON UPDATE NO ACTION
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
