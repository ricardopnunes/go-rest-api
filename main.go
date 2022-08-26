package main

import (
	"fmt"

	"github.com/ricardopnunes/go-rest-api/database"
	"github.com/ricardopnunes/go-rest-api/models"
	"github.com/ricardopnunes/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComDB()
	fmt.Println("Iniciando o Servidor REST!")
	routes.HandleRequest()
}
