package main

import (
	"github.com/carvalhocaio/go-api-gin/models"
	"github.com/carvalhocaio/go-api-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Darth Vader", CPF: "00000000000", RG: "00000000"},
		{Nome: "Leia Organa", CPF: "00000000001", RG: "00000001"},
	}

	routes.HandleRequests()
}
