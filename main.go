package main

import (
	"log"

	"github.com/joao1559/golang-api-scaffold/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	api := api.Server{}
	api.StartServer()
}
