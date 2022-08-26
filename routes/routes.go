package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ricardopnunes/go-rest-api/controllers"
	"github.com/ricardopnunes/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}))(r)))
}
