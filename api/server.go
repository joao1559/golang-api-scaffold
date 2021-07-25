package api

import (
	"log"
	"net/http"
	"os"

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
	// Open connection
	db.InitDb()
	defer db.DBConn.Close()
	err := db.DBConn.Ping()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cors := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedOrigins: []string{"*"},
	})

	var routing = mux.NewRouter()

	healthCheckRepository := repositories.NewMysqlHealthCheckRepository(db.DBConn)
	healthCheckUseCase := usecases.NewHealthCheckUseCase(healthCheckRepository)

	handler.NewHealthCheckHTTPHandler(routing, healthCheckUseCase)

	log.Printf("Listening on port %s...", "4444")
	log.Fatal(http.ListenAndServe(":"+"4444", cors.Handler(routing)))
}
