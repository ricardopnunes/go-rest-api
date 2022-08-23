package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ricardopnunes/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina Inicial")
}

func TodasAsPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
