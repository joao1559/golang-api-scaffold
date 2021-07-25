package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/joao1559/golang-api-scaffold/repositories"
)

func TestAPIHealthCheck(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		a := repositories.NewMysqlHealthCheckRepository(db)
		healthCheck, err := a.Check()
		assert.NoError(t, err)
		assert.NotNil(t, healthCheck)
	})

	t.Run("error", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		db.Close()

		a := repositories.NewMysqlHealthCheckRepository(db)
		healthCheck, err := a.Check()
		assert.NoError(t, err)
		assert.NotNil(t, healthCheck)
		assert.Equal(t, "down", healthCheck.DbUP)
	})
}
