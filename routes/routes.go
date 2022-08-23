package routes

import (
	"log"
	"net/http"

	"github.com/ricardopnunes/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
