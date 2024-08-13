package handler

import (
	"context"
	"net/http"
	"APIsimp/database"
	"APIsimp/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListProdutos(c *gin.Context) {
	
	var produtos []models.Produto

	cursor, err := database.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar produtos"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var produto models.Produto
		if err := cursor.Decode(&produto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar produto"})
			return
		}
		
		produtos = append(produtos, produto)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro durante a iteração"})
		return
	}

	c.JSON(http.StatusOK, produtos)
}
