package main

import (
	"log"

	"github.com/MXkodo/verif/internal/controllers"
	"github.com/MXkodo/verif/internal/repo"
	"github.com/MXkodo/verif/internal/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbRepo, err := repo.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := gin.Default()

	authRepo := repo.New(dbRepo)
	authService := services.NewAuthService(authRepo)
	authController := controllers.New(authService)

	r.POST("/tokens/:user_id", authController.GenerateTokens)
	r.POST("/refresh", authController.RefreshTokens)

	r.Run(":8080")
}
