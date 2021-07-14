package repositories

import (
	"context"

	"github.com/joao1559/golang-api-scaffold/interfaces"
	"github.com/joao1559/golang-api-scaffold/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type mysqlHealthCheckRepository struct {
	Conn *mongo.Client
	Ctx  context.Context
}

// NewMysqlHealthCheckRepository ...
func NewMysqlHealthCheckRepository(Conn *mongo.Client, Ctx context.Context) interfaces.HealthCheckRepository {
	return &mysqlHealthCheckRepository{Conn, Ctx}
}

func (m *mysqlHealthCheckRepository) Check() (*models.HealthCheck, error) {
	a := &models.HealthCheck{}
	dbUp := "up"
	err := m.Conn.Ping(m.Ctx, nil)
	if err != nil {
		dbUp = "down"
	}
	a.DbUP = dbUp
	return a, nil
}
