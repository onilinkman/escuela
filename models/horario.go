package models

//Horario la estructura del horario que cada materia tendra
type Horario struct {
	Idhora      int    `json:"idhora"`
	Horaentrada int    `json:"horaentrada"`
	Horasalida  int    `json:"horasalida"`
	Dia         string `json:"dia"`
	Idmateria   int    `json:"idmateria"`
}

//Horarios Donde se almacena todos los horarios que seran llamados despues
type Horarios []Horario

var queryHorario = `CREATE TABLE if NOT EXISTS horario(
	idhora int primary KEY not null AUTO_INCREMENT,
	horaentrada int not null,
	horasalida int not null,
	dia varchar(15),
	id_materia int,
	CONSTRAINT fk_horario_materia
		FOREIGN KEY (id_materia) REFERENCES materia(idmateria)
)`

//CrearTablaHorario crea la tabla de horario de las materias
func CrearTablaHorario() {
	EjecutarExec(queryHorario)
}
