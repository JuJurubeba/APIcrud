package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetupDatabase(db *mongo.Collection) {
	collection = db
}
