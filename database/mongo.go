package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func SetupDatabase() {
	clientOptions := options.Client().ApplyURI()

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Erro ao verificar a conexão com o MongoDB:", err)
	}

	Collection = client.Database("produtosdb").Collection("produtos")
	log.Println("Conexão com MongoDB estabelecida com sucesso")
}
