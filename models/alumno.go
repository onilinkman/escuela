package models

import "log"

//Alumno crea la estructura de alumno
type Alumno struct {
	Idalumno    int    `json:"idalumno"`
	Nombres     string `json:"nombres"`
	Paterno     string `json:"paterno"`
	Materno     string `json:"materno"`
	Ci          int    `json:"ci"`
	Correo      string `json:"correo"`
	Contrasenia string `json:"contrasenia"`
	Activo      string `json:"activo"`
	Idue        int    `json:"idue"`
}

//Alumnos crea un array de alumnos
type Alumnos []Alumno

var queryAlumnos = `CREATE TABLE if NOT EXISTS alumno(
	idalumno int primary KEY not null AUTO_INCREMENT,
	nombres varchar(45) not null,
	paterno varchar(45) not null,
	materno varchar(45) not null,
	ci int not null,
	correo varchar(45) not null,
	contrasenia varchar(60) not null,
	activo varchar(45) not null,
	ue_idue int,
	CONSTRAINT fk_alumno_unidadEducativa
		FOREIGN KEY (ue_idue) REFERENCES unidadEducativa(idue)
)`

//CrearTablaAlumno crea una tabla de alumnos en la base de datos
func CrearTablaAlumno() {
	EjecutarExec(queryAlumnos)
}

//IngresarAlumno ingresa un alumno a la base de datos
func (al *Alumno) IngresarAlumno() {
	query := `INSERT INTO alumno (nombres,paterno,materno,ci,correo,contrasenia,activo,ue_idue) VALUES (?,?,?,?,?,?,?,?)`
	al.Contrasenia = Encriptar(al.Contrasenia)
	EjecutarExec(query, &al.Nombres, &al.Paterno, &al.Materno, &al.Ci, &al.Correo, &al.Contrasenia, &al.Activo, &al.Idue)
}

//GetAlumno Devuelve un alumno con su id
func GetAlumno(id int) *Alumno {
	query := `SELECT * from alumno WHERE idalumno=?`
	row := EjecutarQuery(query, id)
	alumno := &Alumno{}
	if row.Next() {
		row.Scan(&alumno.Idalumno, &alumno.Nombres, &alumno.Paterno, &alumno.Materno, &alumno.Ci, &alumno.Correo, &alumno.Contrasenia, &alumno.Activo, &alumno.Idue)
	} else {
		log.Println("no se encontro id de alumno en la base de datos")
	}
	return alumno
}

//GetAlumnos devuelve un array de alumnos de la base de datos
func GetAlumnos() Alumnos {
	query := `SELECT * FROM alumno`
	rows := EjecutarQuery(query)
	alumnos := Alumnos{}
	for rows.Next() {
		alumno := Alumno{}
		rows.Scan(&alumno.Idalumno, &alumno.Nombres, &alumno.Paterno, &alumno.Materno, &alumno.Ci, &alumno.Correo, &alumno.Contrasenia, &alumno.Activo, &alumno.Idue)
		alumnos = append(alumnos, alumno)
	}
	return alumnos
}

//GuardarAlumno guarda los cambios que se hicieron al alumno en la base de datos
func (al *Alumno) GuardarAlumno() {
	query := `UPDATE alumno SET nombres=?,paterno=?,materno=?,ci=?,correo=?,contrasenia=?,activo=?,ue_idue=? WHERE idalumno=?`
	al.Contrasenia = Encriptar(al.Contrasenia)
	EjecutarExec(query, &al.Nombres, &al.Paterno, &al.Materno, &al.Ci, &al.Correo, &al.Contrasenia, &al.Activo, &al.Idue, &al.Idalumno)
}

//EliminarAlumno elimina un alumno de la base de datos
func EliminarAlumno(id int) {
	query := `DELETE FROM alumno WHERE idalumno=?`
	EjecutarExec(query, id)
}
