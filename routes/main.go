package routes

import (
	"jwt/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", controllers.Login)
	rg.POST("/register", controllers.Register)
	rg.GET("/acompanhantes", controllers.GetUser)
	rg.GET("/acompanhantes/:genero", controllers.GetUser)
	rg.GET("/acompanhantes/:categoria", controllers.GetUser)
	rg.GET("/acompanhantes/:estado", controllers.GetUser)
	rg.GET("/acompanhantes/:id")
}

func ProtectedRoutes(rg *gin.RouterGroup) {
	rg.GET("/protected", controllers.ProtectedRoute)
	rg.DELETE("/user", controllers.DeleteUser)
	rg.POST("/carro", controllers.AddCarro)
	rg.DELETE("/carro/:id", controllers.RemoverCarro)
}
