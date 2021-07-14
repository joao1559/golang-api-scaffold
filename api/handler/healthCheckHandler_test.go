package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/api/handler"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/usecases/mocks"
)

func TestHealthCheck(t *testing.T) {
	var mockHealthCheck models.HealthCheck

	mockUCase := new(mocks.HealthCheckUseCase)

	mockUCase.On("Check").Return(&mockHealthCheck, nil)

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := handler.HTTPHealthCheckHandler{
		HUsecase: mockUCase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
