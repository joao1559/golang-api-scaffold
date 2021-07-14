package repositories

import (
	"database/sql"

	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/interfaces"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

type mysqlHealthCheckRepository struct {
	Conn *sql.DB
}

// NewMysqlHealthCheckRepository will create an object that represent the article.Repository interface
func NewMysqlHealthCheckRepository(Conn *sql.DB) interfaces.HealthCheckRepository {
	return &mysqlHealthCheckRepository{Conn}
}

func (m *mysqlHealthCheckRepository) Check() (*models.HealthCheck, error) {
	a := &models.HealthCheck{}
	dbUp := "up"
	err := m.Conn.Ping()
	if err != nil {
		dbUp = "down"
	}
	a.DbUP = dbUp
	return a, nil
}
