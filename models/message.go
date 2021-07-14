package models

//Message é a mensagem enviada do client
type Message struct {
	MessageType int    `json:"messageType"`
	Message     string `json:"message"`
}