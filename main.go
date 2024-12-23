package main

import (
	"github.com/carvalhocaio/go-api-gin/database"
	"github.com/carvalhocaio/go-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
