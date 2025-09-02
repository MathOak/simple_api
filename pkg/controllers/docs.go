package controllers

import (
	"encoding/json"
	"net/http"
)

func Docs(w http.ResponseWriter, r *http.Request) {
	docs := map[string]string{
		"GET /items": "Lista itens paginados",
		"GET /docs":  "Documentação das rotas",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}
