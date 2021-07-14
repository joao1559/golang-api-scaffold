package utils_test

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/utils"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
	"testing"
)

func TestRemove(t *testing.T) {
	client := models.Client{}
	var clients []models.Client
	clients = append(clients, client)
	clients = append(clients, client)

	t.Run("success", func(t *testing.T) {
		oldLength := len(clients)

		u := utils.Remove(clients, 0)

		newLength := len(u)
		assert.Equal(t, oldLength-1, newLength)
	})
}

func TestFind(t *testing.T) {
	con := &websocket.Conn{}
	client := models.Client{
		User: "rd_moge",
		Connection: con,
	}
	var clients []models.Client
	clients = append(clients, client)

	t.Run("success", func(t *testing.T) {
		u := utils.Find(clients, con)

		assert.Equal(t, 0, u)
	})

	t.Run("error", func(t *testing.T) {
		con2 := &websocket.Conn{}
		u := utils.Find(clients, con2)

		assert.Equal(t, len(clients), u)
	})
}

func TestFindConn(t *testing.T) {
	con := &websocket.Conn{}
	client := models.Client{
		User: "rd_moge",
		Connection: con,
	}
	var clients []models.Client
	clients = append(clients, client)

	t.Run("success", func(t *testing.T) {
		u := utils.FindConn(clients, "rd_moge")

		assert.Equal(t, 1, len(u))
	})
}

func TestRemoveConn(t *testing.T) {
	con := &websocket.Conn{}
	var connections []*websocket.Conn
	connections = append(connections, con)

	t.Run("success", func(t *testing.T) {
		oldLength := len(connections)
		
		u := utils.RemoveConn(connections, 0)
		newLength := len(u)
		assert.Equal(t, oldLength-1, newLength)
	})
}