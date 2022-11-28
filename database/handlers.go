package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Lemm8/Grupos-CollegeAPI.git/helpers"
	"github.com/Lemm8/Grupos-CollegeAPI.git/validators"
	"github.com/aws/aws-lambda-go/events"
)

const getGruposQuery = `SELECT * FROM Grupo;`

const getGrupoQuery = `SELECT * FROM Grupo WHERE ID = ?;`

const insertGrupoSQL = `INSERT INTO Grupo (Salon, Materia_ID, Docente_ID, Alumno_ID) 
VALUES (?, ?, ?, ?);`

const updateGrupoSQL = `UPDATE Grupo SET Salon = ?, Materia_ID = ?, Docente_ID = ?, Alumno_ID = ? WHERE ID = ?;`

const deleteGrupoSQL = `DELETE FROM Grupo WHERE ID = ?;`

func FetchGrupos(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) ([]*helpers.Grupo, error) {
	rows, err := db.QueryContext(ctx, getGruposQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	grupos := make([]*helpers.Grupo, 0)

	for rows.Next() {
		grupo := &helpers.Grupo{}
		if err := rows.Scan(&grupo.ID, &grupo.Salon, &grupo.Materia_ID,
			&grupo.Docente_ID, &grupo.Alumno_ID); err != nil {
			return nil, err
		}
		grupos = append(grupos, grupo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return grupos, nil
}

func FetchGrupo(ctx context.Context, db *sql.DB, id int) (*helpers.Grupo, error) {
	row := db.QueryRowContext(ctx, getGrupoQuery, id)

	grupo := &helpers.Grupo{}
	if err := row.Scan(&grupo.ID, &grupo.Salon, &grupo.Materia_ID, &grupo.Docente_ID, &grupo.Alumno_ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("No existe un Grupo con el ID %v", id))
		}
		return nil, err
	}
	return grupo, nil
}

func CreateGrupo(ctx context.Context, db *sql.DB, grupo *helpers.Grupo) (*helpers.Grupo, error) {
	if !validators.AlumnoExists(ctx, db, grupo.Alumno_ID) {
		return nil, errors.New(fmt.Sprintf("No existe un alumno con el ID %v", grupo.Alumno_ID))
	}
	if !validators.MateriaExists(ctx, db, grupo.Materia_ID) {
		return nil, errors.New(fmt.Sprintf("No existe una materia con el ID %v", grupo.Materia_ID))
	}
	if !validators.DocenteExists(ctx, db, grupo.Docente_ID) {
		return nil, errors.New(fmt.Sprintf("No existe un docente con el ID %v", grupo.Docente_ID))
	}

	res, err := db.ExecContext(ctx, insertGrupoSQL, grupo.Salon, grupo.Materia_ID,
		grupo.Docente_ID, grupo.Alumno_ID)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	grupo = &helpers.Grupo{
		ID:         int(id),
		Salon:      grupo.Salon,
		Materia_ID: grupo.Materia_ID,
		Docente_ID: grupo.Docente_ID,
		Alumno_ID:  grupo.Alumno_ID,
	}

	return grupo, nil
}

func UpdateGrupo(ctx context.Context, db *sql.DB, id int, grupo *helpers.Grupo) (*helpers.Grupo, error) {

	_, err := FetchGrupo(ctx, db, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No existe un registro con el ID: %v", id))
	}

	if !validators.AlumnoExists(ctx, db, grupo.Alumno_ID) {
		return nil, errors.New(fmt.Sprintf("No existe un alumno con el ID %v", grupo.Alumno_ID))
	}
	if !validators.MateriaExists(ctx, db, grupo.Materia_ID) {
		return nil, errors.New(fmt.Sprintf("No existe una materia con el ID %v", grupo.Materia_ID))
	}
	if !validators.DocenteExists(ctx, db, grupo.Docente_ID) {
		return nil, errors.New(fmt.Sprintf("No existe un docente con el ID %v", grupo.Docente_ID))
	}

	_, err = db.ExecContext(ctx, updateGrupoSQL, grupo.Salon, grupo.Materia_ID,
		grupo.Docente_ID, grupo.Alumno_ID, id)

	if err != nil {
		return nil, err
	}

	updatedGrupo := helpers.Grupo{
		ID:         int(id),
		Salon:      grupo.Salon,
		Docente_ID: grupo.Docente_ID,
		Materia_ID: grupo.Materia_ID,
		Alumno_ID:  grupo.Alumno_ID,
	}

	return &updatedGrupo, nil
}

func DeleteGrupo(ctx context.Context, db *sql.DB, id int) (*helpers.Grupo, error) {
	grupo, err := FetchGrupo(ctx, db, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No existe un registro con el ID: %v", id))
	}

	_, err = db.ExecContext(ctx, deleteGrupoSQL, id)
	if err != nil {
		return nil, err
	}

	return grupo, nil
}
