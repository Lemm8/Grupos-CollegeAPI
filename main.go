package main

import (
	"context"
	"database/sql"

	"github.com/Lemm8/Grupos-CollegeAPI.git/api"
	"github.com/Lemm8/Grupos-CollegeAPI.git/database"
	"github.com/Lemm8/Grupos-CollegeAPI.git/helpers"
	"github.com/Lemm8/Grupos-CollegeAPI.git/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var db *sql.DB

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	if !validators.IsValidPath(event.Path) {
		return helpers.ServerError(500, "Invalid Path"), nil
	}

	// CONNECT TO DB
	dbConnection, err := database.GetConnection()
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}
	db = dbConnection

	switch event.HTTPMethod {
	case "GET":
		apiResponse, err := api.GetGrupos(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiResponse, nil

	case "POST":
		apiResponse, err := api.PostGrupo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiResponse, nil

	case "PUT":
		apiResponse, err := api.PutGrupo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiResponse, nil

	case "DELETE":
		apiResponse, err := api.DeleteGrupo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiResponse, nil

	default:
		return helpers.UnhandledMethod(), nil
	}
}

func main() {
	lambda.Start(handler)
}
