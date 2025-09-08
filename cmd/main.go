package main

import (
	"simple_api/pkg/config"
	database "simple_api/pkg/database/mysql"
	"simple_api/pkg/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	db := database.Connect()
	println("DB connected successfully.", db)

	r := gin.Default()
	r.Use(DBMiddleware(db))
	routes.RegisterRoutes(r)
	println("Server running on port", config.AppConfig.Port)
	r.Run(":" + config.AppConfig.Port)
}
