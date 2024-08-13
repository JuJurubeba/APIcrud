package handler

import (
	"APIsimp/database"
	"APIsimp/models"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduto(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		log.Printf("Erro ao fazer o bind do JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	produto.ID = primitive.NewObjectID()
	_, err := database.Collection.InsertOne(context.TODO(), produto)
	if err != nil {
		log.Printf("Erro ao inserir o produto no banco de dados: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, produto)
}
