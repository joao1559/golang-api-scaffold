package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joao1559/golang-api-scaffold/config"
	"github.com/joao1559/golang-api-scaffold/interfaces"
)

//HTTPNoteHandler represent the httphandler for notes
type HTTPNoteHandler struct {
	NUseCase interfaces.NoteUseCase
}

//NewNoteHTTPHandler ...
func NewNoteHTTPHandler(e *mux.Router, us interfaces.NoteUseCase) {
	handler := &HTTPNoteHandler{
		NUseCase: us,
	}
	s := e.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/notes", handler.Get).Methods(http.MethodGet)
	s.HandleFunc("/note", handler.Insert).Methods(http.MethodPost)
}

//Get handler da rota GET /notes
func (n *HTTPNoteHandler) Get(w http.ResponseWriter, req *http.Request) {
	var t config.Config
	notes, err := n.NUseCase.Get()
	if err != nil {
		t.RespondWithError(w, 500, "INTERNAL ERROR", err.Error())
		return
	}

	t.ResponseWithJSON(w, http.StatusOK, notes)
}

//Insert handler da rota POST /notes
func (n *HTTPNoteHandler) Insert(w http.ResponseWriter, req *http.Request) {
	var t config.Config

	body, _ := ioutil.ReadAll(req.Body)
	err := n.NUseCase.Insert(body)
	if err != nil {
		t.RespondWithError(w, 500, "INTERNAL ERROR", err.Error())
		return
	}

	t.ResponseWithJSON(w, http.StatusCreated, nil)
}
