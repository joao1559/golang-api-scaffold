package config

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/sirupsen/logrus"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

//Config o aplicativo
type Config struct {
	Env string
}

//RespondWithError corresponde a funão que restorna erro
func (a *Config) RespondWithError(w http.ResponseWriter, code int, message string, moreInfo string) {
	var m models.ResponseError
	m.DeveloperMessage = message
	m.UserMessage = "Erro"
	m.ErrorCode = code
	m.MoreInfo = moreInfo
	respondWithError(w, code, m)
}

//RespondWithErrorNew corresponde a funão que restorna erro
func (a *Config) RespondWithErrorNew(w http.ResponseWriter, code int, codeError int) {
	m := getMessageError(codeError)
	respondWithError(w, code, m)
}

func respondWithError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//ResponseWithJSON corresponde ao metódo que retorna sucesso
func (a *Config) ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	var r models.ResponseSuccess
	r.Records = payload
	lenPayload := reflect.ValueOf(payload)
	r.Meta.RecordCount = 1
	r.Meta.Limit = 1
	if lenPayload.Kind() == reflect.Slice {
		r.Meta.Limit = lenPayload.Len()
		r.Meta.RecordCount = lenPayload.Len()
	}
	response, _ := json.Marshal(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//GetStatusCode ...
func (a *Config) GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case sql.ErrNoRows:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

//GetMessageError ...
func getMessageError(errorCode int) *models.ResponseError {
	var m models.ResponseError
	switch errorCode {
	case 10000:
		return &models.ResponseError{
			DeveloperMessage: "Internal server error {0}",
			UserMessage:      "Was encountered an error when processing your request. We apologize for the inconvenience",
			MoreInfo:         "http://www.developer.apiluiza.com.br",
			ErrorCode:        10000,
		}
	case 20002:
		m.DeveloperMessage = "Resource not found"
		m.UserMessage = "Resource not found"
		m.MoreInfo = "http://www.developer.apiluiza.com.br"
		m.ErrorCode = 20002
		return &m
	default:
		return &models.ResponseError{
			DeveloperMessage: "Resource not found",
			UserMessage:      "Resource not found",
			MoreInfo:         "http://www.developer.apiluiza.com.br",
			ErrorCode:        404,
		}
	}
}
