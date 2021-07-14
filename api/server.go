package api

import (
	"log"
	"net/http"
	"os"

	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/api/handler"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/repositories"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/usecases"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/config/db"
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

	notificationRepository := repositories.NewMySQLNotification(db.DBConn)
	notificationUseCase := usecases.NewNotificationUseCase(notificationRepository)
	healthCheckRepository := repositories.NewMysqlHealthCheckRepository(db.DBConn)
	healthCheckUseCase := usecases.NewHealthCheckUseCase(healthCheckRepository)

	handler.NewNotificationHTTPHandler(routing, notificationUseCase)
	handler.NewHealthCheckHTTPHandler(routing, healthCheckUseCase)

	log.Printf("Listening on port %s...", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), cors.Handler(routing)))
}
