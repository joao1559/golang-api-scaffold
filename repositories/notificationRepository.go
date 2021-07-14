package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/interfaces"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

type mySQLNotification struct {
	Conn *sql.DB
}

//NewMySQLNotification retorna a instancia do repository
func NewMySQLNotification(Conn *sql.DB) interfaces.NotificationRepository {
	return &mySQLNotification{Conn}
}

//GetAll busca todas as notificações que não foram visualizadas
func (m *mySQLNotification) GetAll(not *models.Notification) ([]*models.Notification, error) {
	var values []interface{}
	var where []string

	if not.Usuario != "" {
		where = append(where, "usuario = ?")
		values = append(values, not.Usuario)
	} else {
		where = append(where, "data_envio IS NULL")
	}

	rows, err := m.Conn.Query(`SELECT id, titulo_notificacao, mensagem, link_notificacao, usuario, 
									data_envio, data_visualizacao, data_cadastro, usuario_cadastro, data_alteracao, usuario_alteracao
									FROM mag_t_tp_notificacao
							   WHERE `+strings.Join(where, " AND ")+` ORDER BY data_envio DESC`, values...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Notification, 0)

	for rows.Next() {
		a := &models.Notification{}

		var (
			id                sql.NullInt64
			tituloNotificacao sql.NullString
			mensagem          sql.NullString
			linkNotificacao   sql.NullString
			usuario           sql.NullString
			dataEnvio         sql.NullString
			dataVisualizacao  sql.NullString
			dataCadastro      sql.NullString
			usuarioCadastro   sql.NullString
			dataAlteracao     sql.NullString
			usuarioAlteracao  sql.NullString
		)

		err := rows.Scan(
			&id,
			&tituloNotificacao,
			&mensagem,
			&linkNotificacao,
			&usuario,
			&dataEnvio,
			&dataVisualizacao,
			&dataCadastro,
			&usuarioCadastro,
			&dataAlteracao,
			&usuarioAlteracao,
		)

		if err != nil {
			return nil, err
		}

		a.ID = id.Int64
		a.Titulo = tituloNotificacao.String
		a.Mensagem = mensagem.String
		a.Link = linkNotificacao.String
		a.Usuario = usuario.String
		a.DataEnvio = dataEnvio.String
		a.DataVisualizacao = dataVisualizacao.String
		a.DataCadastro = dataCadastro.String
		a.UsuarioCadastro = usuarioCadastro.String
		a.DataAlteracao = dataAlteracao.String
		a.UsuarioAlteracao = usuarioAlteracao.String

		result = append(result, a)
	}

	return result, nil
}

func (m *mySQLNotification) Update(not *models.Notification) error {

	query := `UPDATE mag_t_tp_notificacao
				 SET data_envio 		= ?,
					 titulo_notificacao = ?,
					 mensagem			= ?, 
					 link_notificacao	= ?, 
					 usuario			= ?,
					 data_alteracao  	= now(),
					 usuario_alteracao 	= "TIO-PATINHAS-NOTIFICACAO-API"
	           WHERE id             = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil
	}

	res, err := stmt.Exec(not.DataEnvio, not.Titulo, not.Mensagem, not.Link, not.Usuario, not.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return err
	}
	return nil
}

func (m *mySQLNotification) Read(not *models.Notification) error {

	query := `UPDATE mag_t_tp_notificacao
				SET data_visualizacao	= now(),
					data_alteracao  	= now(),
					usuario_alteracao 	= "TIO-PATINHAS-NOTIFICACAO-API"
  			  WHERE id = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil
	}

	res, err := stmt.Exec(not.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return err
	}
	return nil
}
