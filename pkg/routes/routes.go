package routes

import (
	"net/http"
	"simple_api/pkg/controllers"

	"gorm.io/gorm"
)

func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		controllers.Docs(w, r)
	})
	// Exemplo de rota paginada
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetItems(w, r, db)
	})
}
