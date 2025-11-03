package main

import (
	"github.com/DevCarlosOli/gin-api-rest/database"
	"github.com/DevCarlosOli/gin-api-rest/models"
	"github.com/DevCarlosOli/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Carlos Filho", CPF: "00000000000", RG: "123456"},
		{Nome: "Aline Ferreira", CPF: "99999999999", RG: "654321"},
	}
	routes.HandleRequests()
}
