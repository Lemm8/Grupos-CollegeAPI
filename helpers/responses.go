package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func UnhandledMethod() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusBadRequest,
		Message: "Unhandled method, try again",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       string(body),
	}
}

func ServerError(status int, errMessage string) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  status,
		Message: errMessage,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(body),
	}
}

func ServerSuccess(status int, msg string) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: msg,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func GetGruposResponse(grupos []*Grupo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&ListGruposResponse{
		Status: http.StatusOK,
		Grupos: grupos,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func GetGrupoResponse(grupo *Grupo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&GrupoResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Grupo:   grupo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func PostGrupoResponse(grupo *Grupo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&GrupoResponse{
		Status:  http.StatusOK,
		Message: "Grupo Created",
		Grupo:   grupo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func PutGrupoResponse(grupo *Grupo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&GrupoResponse{
		Status:  http.StatusOK,
		Message: "Grupo Updated",
		Grupo:   grupo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func DeleteGrupoResponse(grupo *Grupo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&GrupoResponse{
		Status:  http.StatusOK,
		Message: "Grupo Deleted",
		Grupo:   grupo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}
