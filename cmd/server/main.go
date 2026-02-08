package main

import (
	"server/server/internal/config"
	"server/server/internal/database"
	"server/server/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()

r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"}, // DEV ONLY
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Content-Type", "Authorization"},
    AllowCredentials: false, // HARUS false kalau origin "*"
}))

routes.RegisterRoutes(r)

// listen ke semua network
r.Run("0.0.0.0:8080")

}
