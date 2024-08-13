package handler

import (
    "context"
    "log"
    "net/http"
    "APIsimp/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProduto(c *gin.Context) {
    id := c.Param("id")

    
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Println("Erro ao converter ID:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    filter := bson.M{"_id": objectId}


    result, err := database.Collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Println("Erro ao deletar produto:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto"})
        return
    }

    if result.DeletedCount == 0 {
        log.Println("Produto não encontrado")
        c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
        return
    }

    log.Println("Produto deletado com sucesso")
    c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}
