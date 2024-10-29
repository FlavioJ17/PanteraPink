package controllers

import (
	"jwt/initializers"
	"jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCarro(c *gin.Context) {
	var body struct {
		Modelo     string `json:"modelo"`
		Fabricante string `json:"fabricante"`
		Ano        string `json:"ano"`
		Cor        string `json:"cor"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao enviar"})
		return
	}

	if body.Modelo == "" || body.Fabricante == "" || body.Ano == "" || body.Cor == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos vazio"})
	}

	carro := models.Carro{
		Modelo:     body.Modelo,
		Fabricante: body.Fabricante,
		Ano:        body.Ano,
		Cor:        body.Cor,
	}

	result := initializers.DB.Create(&carro)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falhou ao criar o modelo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Carro created successfully"})
}

func RemoverCarro(c *gin.Context) {
	id := c.Param("id")
	_, exists := c.Get("username")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	var carro models.Carro
	result := initializers.DB.First(&carro, "id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := initializers.DB.Delete(&carro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falhou ao deletar o usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sucesso"})

}
