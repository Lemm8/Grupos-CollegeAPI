package validators

import (
	"context"
	"database/sql"

	"github.com/Lemm8/Grupos-CollegeAPI.git/helpers"
)

const getAlumnoQuery = `SELECT * FROM Alumno WHERE ID = ?;`
const getMateriaQuery = `SELECT * FROM Materia WHERE ID = ?;`
const getDocenteQuery = `SELECT * FROM Docente WHERE ID = ?;`

func IsValidPath(path string) bool {
	if path != "/Grupos" {
		return false
	}

	return true
}

func AlumnoExists(ctx context.Context, db *sql.DB, id int) bool {
	row := db.QueryRowContext(ctx, getAlumnoQuery, id)

	alumno := &helpers.Alumno{}
	if err := row.Scan(&alumno.ID, &alumno.Nombre, &alumno.Apellido, &alumno.Matricula,
		&alumno.Fecha_Nacimiento, &alumno.Semestre, &alumno.Carreras_ID); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}

func DocenteExists(ctx context.Context, db *sql.DB, id int) bool {
	// QUERY CALIFICACION BY ID
	row := db.QueryRowContext(ctx, getDocenteQuery, id)

	docente := &helpers.Docente{}
	if err := row.Scan(&docente.ID, &docente.Nombre, &docente.Apellido, &docente.Matricula, &docente.Fecha_Nacimiento, &docente.Titulo, &docente.Correo, &docente.Telefono); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}

func MateriaExists(ctx context.Context, db *sql.DB, id int) bool {
	// QUERY CALIFICACION BY ID
	row := db.QueryRowContext(ctx, getMateriaQuery, id)

	materia := &helpers.Materia{}
	if err := row.Scan(&materia.ID, &materia.Nombre, &materia.IsTroncoComun); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}
