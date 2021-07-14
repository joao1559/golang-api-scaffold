package models

import "github.com/gorilla/websocket"

//Client é o websocket conectado
type Client struct {
	Connection *websocket.Conn
	User       string
	ID         int
}