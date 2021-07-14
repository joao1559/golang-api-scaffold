package utils

import (
	"github.com/gorilla/websocket"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

//Remove é uma função usada para remover um client de conexão do array
func Remove(s []models.Client, i int) []models.Client {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

//RemoveConn é uma função para remover a Connection do socket de um array
func RemoveConn(s []*websocket.Conn, i int) []*websocket.Conn {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

//Find é uma funçãp para encontrar um Client pela connection
func Find(a []models.Client, x *websocket.Conn) int {
	for i, n := range a {
		if x == n.Connection {
			return i
		}
	}
	return len(a)
}

//FindConn é uma função que retorna todas as connections de um usuario
func FindConn(a []models.Client, x string) []*websocket.Conn {
	var conns []*websocket.Conn

	for _, n := range a {
		if x == n.User {
			conns = append(conns, n.Connection)
		}
	}
	return conns
}