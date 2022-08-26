package main

import (
	"fmt"

	"github.com/ricardopnunes/go-rest-api/database"
	"github.com/ricardopnunes/go-rest-api/routes"
)

func main() {
	database.ConectaComDB()
	fmt.Println("Iniciando o Servidor REST!")
	routes.HandleRequest()
}
