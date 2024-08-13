package main

import (

	 "APIsimp/database"
    "APIsimp/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	database.SetupDatabase()

	r := gin.Default()

	
	r.POST("/produtos", handler.CreateProduto)
	r.GET("/produtos", handler.ListProdutos)
	r.PUT("/produtos/:id", handler.UpdateProduto)
	r.DELETE("/produtos/:id", handler.DeleteProduto)

	r.Run(":8080") 
	r.SetTrustedProxies(nil)

}
