package config

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/joao1559/golang-api-scaffold/models"
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
