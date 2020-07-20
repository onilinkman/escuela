package models

//Requisito modelo de la estructura requisito
type Requisito struct {
	Idmateria    int `json:"idmateria"`
	Idmateriareq int `json:"idmateriareq"`
}

var query = `CREATE TABLE if NOT EXISTS requisito(
	id_materia int,
	CONSTRAINT fk_requisito_materia
		FOREIGN KEY (id_materia) REFERENCES materia(idmateria),
	id_materiareq int,
	CONSTRAINT fk_requisito2_materia
		FOREIGN KEY (id_materiareq) REFERENCES materia(idmateria)
)`

//CrearTablaRequisito crea la tabla requisito en la base de datos
func CrearTablaRequisito() {
	EjecutarExec(query)
}

//AgregarRequisito se agrega el requisito que necesita una materia a la base de datos
func (req *Requisito) AgregarRequisito() {
	query := `INSERT INTO requisito (id_materia,id_materiareq) VALUES (?,?)`
	EjecutarExec(query, &req.Idmateria, &req.Idmateriareq)
}

//GetRequisitosMateria obtiene las materias que necesita para que se habilite el siguiente curso
func GetRequisitosMateria(idnat int) Materias {
	query := `SELECT ma.* FROM materia ma,(SELECT id_materiareq FROM requisito WHERE id_materia=?) tmp
	WHERE tmp.id_materiareq=ma.idmateria`
	rows := EjecutarQuery(query, idnat)
	materias := Materias{}
	for rows.Next() {
		materia := Materia{}
		rows.Scan(&materia.Idmateria, &materia.Notamin, &materia.Notafinal, &materia.Nombre, &materia.Gestion)
		materias = append(materias, materia)
	}
	return materias

}
