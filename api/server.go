package api

import (
	"context"
	"log"
	"net/http"

	"github.com/joao1559/golang-api-scaffold/api/handler"
	"github.com/joao1559/golang-api-scaffold/repositories"
	"github.com/joao1559/golang-api-scaffold/usecases"

	"github.com/gorilla/mux"
	"github.com/joao1559/golang-api-scaffold/config/db"
	"github.com/rs/cors"
)

//Server ...
type Server struct{}

//StartServer inicia o servidor
func (s *Server) StartServer() {
	var ctx = context.TODO()
	// Open connection
	DBConn := db.InitDb(ctx)
	defer DBConn.Disconnect(ctx)

	cors := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedOrigins: []string{"*"},
	})

	var routing = mux.NewRouter()

	healthCheckRepository := repositories.NewMysqlHealthCheckRepository(DBConn, ctx)
	healthCheckUseCase := usecases.NewHealthCheckUseCase(healthCheckRepository)

	noteRepository := repositories.NewMongoNoteRepository(DBConn, ctx)
	noteUseCase := usecases.NewNoteUseCase(noteRepository)

	handler.NewHealthCheckHTTPHandler(routing, healthCheckUseCase)
	handler.NewNoteHTTPHandler(routing, noteUseCase)

	log.Printf("Listening on port %s...", "4444")
	log.Fatal(http.ListenAndServe(":"+"4444", cors.Handler(routing)))
}
