package handler_test

import (
	"strconv"
	"github.com/stretchr/testify/mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/api/handler"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/usecases/mocks"
)

func TestGetAll(t *testing.T) {
	mockNotification := &models.Notification{}

	mockUCase := new(mocks.NotificationUseCase)

	mockListNotification := make([]*models.Notification, 0)
	mockListNotification = append(mockListNotification, mockNotification)

	mockUCase.On("GetAll", mockNotification).Return(mockListNotification, nil)

	req, err := http.NewRequest("GET", "/notifications", strings.NewReader(""))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := handler.HTTPNotificationHandler{
		NotificationUseCase: mockUCase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/notifications", handler.GetAll).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUCase.AssertExpectations(t)
}

func TestRead(t *testing.T) {
	mockNotification := models.Notification{}
	tempMockNotification := mockNotification
	mockUCase := new(mocks.NotificationUseCase)

	j, err := json.Marshal(tempMockNotification)
	assert.NoError(t, err)

	notificationId := mockNotification.ID

	mockUCase.On("Read", mock.AnythingOfType("*models.Notification")).Return(nil)

	req, err := http.NewRequest("PUT", "/notification/"+strconv.Itoa(int(notificationId)) + "/read", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := handler.HTTPNotificationHandler{
		NotificationUseCase: mockUCase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/notification/{id:[0-9]+}/read", handler.Read).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUCase.AssertExpectations(t)
}