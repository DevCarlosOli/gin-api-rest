package main

import (
	"github.com/DevCarlosOli/gin-api-rest/database"
	"github.com/DevCarlosOli/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
