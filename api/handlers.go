package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/Lemm8/Grupos-CollegeAPI.git/database"
	"github.com/Lemm8/Grupos-CollegeAPI.git/helpers"
	"github.com/aws/aws-lambda-go/events"
)

func GetGrupos(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		grupo := &helpers.Grupo{}
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}

		grupo, err = database.FetchGrupo(ctx, db, intID)
		return helpers.GetGrupoResponse(grupo), nil
	}

	alumnosInfo, err := database.FetchGrupos(ctx, db, req)
	if err != nil {
		return nil, err
	}

	return helpers.GetGruposResponse(alumnosInfo), nil

}

func PostGrupo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	grupo := &helpers.Grupo{}
	err := json.Unmarshal([]byte(req.Body), &grupo)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	grupo, err = database.CreateGrupo(ctx, db, grupo)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	return helpers.PostGrupoResponse(grupo), nil
}

func PutGrupo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	grupo := &helpers.Grupo{}
	err := json.Unmarshal([]byte(req.Body), &grupo)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}
		updatedGrupo, err := database.UpdateGrupo(ctx, db, intID, grupo)
		if err != nil {
			return nil, err
		}
		return helpers.PutGrupoResponse(updatedGrupo), nil
	}

	return nil, errors.New("se debe incluir el id del alumno_info para actualizarlo")
}

func DeleteGrupo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// VALIDATE ID EXISTS IN QUERY
	id := req.QueryStringParameters["id"]
	if len(id) < 1 {
		return helpers.ServerError(400, "Se debe incluir el ID"), nil
	}

	// CONVERT ID TO INT
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("el id no es válido")
	}

	// DELETE ALUMNO
	grupo := &helpers.Grupo{}

	grupo, err = database.DeleteGrupo(ctx, db, intID)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	return helpers.DeleteGrupoResponse(grupo), nil
}
