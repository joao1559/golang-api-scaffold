package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joao1559/golang-api-scaffold/api/handler"
	"github.com/joao1559/golang-api-scaffold/models"
	"github.com/joao1559/golang-api-scaffold/usecases/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	var mockHealthCheck models.HealthCheck

	mockUCase := new(mocks.HealthCheckUseCase)

	mockUCase.On("Check").Return(&mockHealthCheck, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()

	handler.NewHealthCheckHTTPHandler(router, mockUCase)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
