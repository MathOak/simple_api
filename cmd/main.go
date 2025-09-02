package main

import (
	"net/http"
	database "simple_api/pkg/database/mysql"
	"simple_api/pkg/routes"
)

func main() {
	db := database.Connect()
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, db)
	http.ListenAndServe(":8080", mux)
}
