package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/config"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/interfaces"
)

//HTTPHealthCheckHandler represent the httphandler for article
type HTTPHealthCheckHandler struct {
	HUsecase interfaces.HealthCheckUseCase
}

//NewHealthCheckHTTPHandler ...
func NewHealthCheckHTTPHandler(e *mux.Router, us interfaces.HealthCheckUseCase) {
	handler := &HTTPHealthCheckHandler{
		HUsecase: us,
	}
	s := e.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
}

// HealthCheck handler da rota GET /health
func (a *HTTPHealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	var t config.Config
	h, _ := a.HUsecase.Check()
	h.Status = "up"
	t.ResponseWithJSON(w, http.StatusOK, h)
}
