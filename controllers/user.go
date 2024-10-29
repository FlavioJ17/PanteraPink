package controllers

import (
	"jwt/initializers"
	"jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	username, exists := c.Get("username")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falhou ao deletar o usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sucesso"})

}
func GetUser(c *gin.Context) {

	var users []models.User

	// Validação do corpo da requisição
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, users)

}
