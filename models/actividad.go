package models

import "log"

//Actividad la estructura de la actividad
type Actividad struct {
	Idactividad int     `json:"idactividad"`
	Fecha       string  `json:"fecha"`
	Fechavenc   string  `json:"fechavenc"`
	Nombre      string  `json:"nombre"`
	Tipo        string  `json:"tipo"`
	Nota        float32 `json:"nota"`
}

//Actividades array de actividades
type Actividades []Actividad

var queryActividad = `CREATE TABLE if NOT EXISTS actividad(
	idactividad int primary KEY not null AUTO_INCREMENT,
	fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	fechavenc TIMESTAMP not null,
	nombre varchar(45) not null,
	tipo varchar(25) not null,
	nota float not null
)`

//CrearTablaActividad se crea la tabla de actividades
func CrearTablaActividad() {
	EjecutarExec(queryActividad)
}

//IngresarActividad ingresa una actividad a la base de datos
func (ac *Actividad) IngresarActividad() {
	query := `INSERT INTO actividad (fechavenc,nombre,tipo,nota) VALUES (?,?,?,?)`
	EjecutarExec(query, &ac.Fechavenc, &ac.Nombre, &ac.Tipo, &ac.Nota)
}

//GetActividad devuelve una actividad de la base de datos segun su id
func GetActividad(idact int) Actividad {
	query := `SELECT * FROM actividad WHERE idactividad=?`
	row := EjecutarQuery(query, idact)
	actividad := Actividad{}
	if row.Next() {
		row.Scan(&actividad.Idactividad, &actividad.Fecha, &actividad.Fechavenc, &actividad.Nombre, &actividad.Tipo, &actividad.Nota)
	} else {
		log.Println("no se encontro la actividad con esa ci")
	}
	return actividad
}

//GetActividades recupera un array de actividad
func GetActividades() Actividades {
	query := `SELECT * FROM actividad`
	rows := EjecutarQuery(query)
	actividades := Actividades{}
	for rows.Next() {
		actividad := Actividad{}
		rows.Scan(&actividad.Idactividad, &actividad.Fecha, &actividad.Fechavenc, &actividad.Nombre, &actividad.Tipo, &actividad.Nota)
		actividades = append(actividades, actividad)
	}
	return actividades
}

//GuardarActividad guarda los cambios de una actividad
func (ac *Actividad) GuardarActividad() {
	query := `UPDATE actividad SET fechavenc=?,nombre=?,tipo=?,nota=? WHERE idactividad=?`
	EjecutarExec(query, &ac.Fechavenc, &ac.Nombre, &ac.Tipo, &ac.Nota, &ac.Idactividad)
}

//EliminarActividad elimina una actividad depende de su id
func EliminarActividad(id int) {
	query := `DELETE FROM actividad WHERE idactividad=?`
	EjecutarExec(query, id)
}
