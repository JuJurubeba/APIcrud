package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Produto representa a estrutura na base de dados.
type Produto struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Nome      string             `json:"nome" bson:"nome"`
    Preco     float64            `json:"preco" bson:"preco"`
    Quantidade int               `json:"quantidade" bson:"quantidade"`
}
