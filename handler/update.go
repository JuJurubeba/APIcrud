package handler

import (
    "context"
    "log"
    "net/http"
    "APIsimp/models"
    "APIsimp/database"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProduto(c *gin.Context) {
    id := c.Param("id")

    // Verifica se o ID tem 24 caracteres hexadecimais
    if len(id) != 24 {
        log.Println("ID inválido: comprimento incorreto")
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Println("Erro ao converter ID:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var produto models.Produto
    if err := c.ShouldBindJSON(&produto); err != nil {
        log.Println("Erro ao bindar JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    update := bson.M{
        "$set": bson.M{
            "nome":       produto.Nome,
            "preco":      produto.Preco,
            "quantidade": produto.Quantidade,
        },
    }

    result, err := database.Collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
    if err != nil {
        log.Println("Erro ao atualizar produto:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"})
        return
    }

    if result.MatchedCount == 0 {
        log.Println("Produto não encontrado")
        c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
        return
    }

    log.Println("Produto atualizado com sucesso")
    produto.ID = objID
    c.JSON(http.StatusOK, produto)
}
