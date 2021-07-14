package main

import (
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/api"
	"log"

)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	api := api.Server{}
	api.StartServer()
}