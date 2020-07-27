package models

import "log"

//Profesor modelo de la clase profesor
type Profesor struct {
	Ci          int    `json:"ci"`
	Nombres     string `json:"nombres"`
	Paterno     string `json:"paterno"`
	Materno     string `json:"materno"`
	Correo      string `json:"correo"`
	Contrasenia string `json:"contrasenia"`
	Idue        int    `json:"idue"`
}

//Profesores array de profesores
type Profesores []Profesor

var queryProfe = `CREATE TABLE if NOT EXISTS profesor(
	ci int primary KEY not null,
	nombres varchar(45) not null,
	paterno varchar(45) not null,
	materno varchar(45) not null,
	correo varchar(45) not null,
	contrasenia varchar(60) not null,
	ue_idue int,
	CONSTRAINT fk_profesor_unidadEducativa
		FOREIGN KEY (ue_idue) REFERENCES unidadEducativa(idue)
)`

//CrearTablaProfesor crea la tabla profesor en la base de datos si no existe
func CrearTablaProfesor() {
	EjecutarExec(queryProfe)
}

//IngresarProfesor ingresa datos de un profesor en la tabla profesor
func (p *Profesor) IngresarProfesor() {
	query := `INSERT INTO profesor VALUES (?,?,?,?,?,?,?)`
	p.Contrasenia = Encriptar(p.Contrasenia)
	EjecutarExec(query, &p.Ci, &p.Nombres, &p.Paterno, &p.Materno, p.Correo, p.Contrasenia, p.Idue)
}

//GetProfesores devuelve todos los profesores
func GetProfesores() Profesores {
	profesores := Profesores{}
	query := `select * from profesor`
	rows := EjecutarQuery(query)
	for rows.Next() {
		profesor := Profesor{}
		rows.Scan(&profesor.Ci, &profesor.Nombres, &profesor.Paterno, &profesor.Materno, &profesor.Correo, &profesor.Contrasenia, &profesor.Idue)
		profesores = append(profesores, profesor)
	}
	return profesores

}

//GetProfesor obtiene un profesor con su ci
func GetProfesor(cis int) *Profesor {
	query := `select * from profesor WHERE ci=?`
	row := EjecutarQuery(query, cis)
	profesor := &Profesor{}
	if row.Next() {
		row.Scan(&profesor.Ci, &profesor.Nombres, &profesor.Paterno, &profesor.Materno, &profesor.Correo, &profesor.Contrasenia, &profesor.Idue)
	} else {
		log.Println("no se encontro ningun profesor con esa ci")
	}
	return profesor
}

//GuardarProfesor guarda los datos del profesor que fueron cambiados
func (p *Profesor) GuardarProfesor() {
	query := `UPDATE profesor SET nombres=?,paterno=?,materno=?,correo=?,contrasenia=?,ue_idue=? WHERE ci=?`
	p.Contrasenia = Encriptar(p.Contrasenia)
	EjecutarExec(query, &p.Nombres, &p.Paterno, &p.Materno, &p.Correo, &p.Contrasenia, &p.Idue, p.Ci)
}

//EliminarProfesor elimina un profesor por su ci
func EliminarProfesor(ci int) {
	query := `DELETE FROM profesor WHERE ci=?`
	EjecutarExec(query, ci)
}
