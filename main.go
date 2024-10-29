package main

import (
	"jwt/initializers"
	"jwt/middleware"
	"jwt/models"
	"jwt/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Conecta ao banco de dados
	initializers.ConnectToDB()

	// Migra o modelo de usuário
	initializers.DB.Migrator().AutoMigrate([]interface{}{&models.User{}, models.Carro{}}...)
}

func main() {
	r := gin.Default()

	// Rotas públicas
	public := r.Group("/api")
	routes.PublicRoutes(public)

	// Rotas protegidas por JWT
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware()) // Middleware de autenticação
	routes.ProtectedRoutes(protected)

	r.Run(":8080") // Inicia o servidor na porta 8080
}
