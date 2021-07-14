package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/repositories"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mockNotification = &models.Notification{
	ID:               1,
	Titulo:           "Notificacao",
	Mensagem:         "Nova notificacao",
	Link:             "http://localhost:8000/",
	Usuario:          "rd_moge",
	DataEnvio:        "",
	DataVisualizacao: "",
	DataCadastro:     "2019-10-01",
	UsuarioCadastro:  "rd_moge",
	DataAlteracao:    "",
	UsuarioAlteracao: "",
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "titulo_notificacao", "mensagem", "link_notificacao",
		"usuario", "data_envio", "data_visualizacao", "data_cadastro", "usuario_cadastro", "data_alteracao", "usuario_alteracao"}).
		AddRow(
			mockNotification.ID,
			mockNotification.Titulo,
			mockNotification.Mensagem,
			mockNotification.Link,
			mockNotification.Usuario,
			mockNotification.DataEnvio,
			mockNotification.DataVisualizacao,
			mockNotification.DataCadastro,
			mockNotification.UsuarioCadastro,
			mockNotification.DataAlteracao,
			mockNotification.UsuarioAlteracao,
		)

	query := `SELECT id, titulo_notificacao, mensagem, link_notificacao, usuario, 
				data_envio, data_visualizacao, data_cadastro, usuario_cadastro, data_alteracao, usuario_alteracao
				FROM mag_t_tp_notificacao
			WHERE usuario = ? ORDER BY data_envio DESC`

	mock.ExpectQuery(query).WithArgs(mockNotification.Usuario).WillReturnRows(rows)
	repository := repositories.NewMySQLNotification(db)

	_, err = repository.GetAll(mockNotification)
	assert.NoError(t, err)

	mockNotification.Usuario = ""

	query = `SELECT id, titulo_notificacao, mensagem, link_notificacao, usuario, 
				data_envio, data_visualizacao, data_cadastro, usuario_cadastro, data_alteracao, usuario_alteracao
				FROM mag_t_tp_notificacao
			WHERE data_envio IS NULL ORDER BY data_envio DESC`

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)
	repository = repositories.NewMySQLNotification(db)

	_, err = repository.GetAll(mockNotification)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	notification := mockNotification

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	query := `UPDATE mag_t_tp_notificacao
				SET data_envio 			= ?,
					titulo_notificacao 	= ?,
					mensagem			= ?, 
					link_notificacao	= ?, 
					usuario				= ?,
					data_alteracao  	= now(),
					usuario_alteracao 	= "TIO-PATINHAS-NOTIFICACAO-API"
			  WHERE id = ?`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(notification.DataEnvio, notification.Titulo, notification.Mensagem,
		notification.Link, notification.Usuario, notification.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	rel := repositories.NewMySQLNotification(db)

	err = rel.Update(notification)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	notification := mockNotification

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	query := `UPDATE mag_t_tp_notificacao
				SET data_visualizacao	= now(),
					data_alteracao  	= now(),
					usuario_alteracao 	= "TIO-PATINHAS-NOTIFICACAO-API"
			  WHERE id = ?`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(notification.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	rel := repositories.NewMySQLNotification(db)

	err = rel.Read(notification)
	assert.NoError(t, err)
}
