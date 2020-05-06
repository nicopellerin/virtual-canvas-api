package database

import (
	"context"
	"fmt"
	"log"

	"github.com/nicopellerin/virtual-canvas-api/graph/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "virtualcanvas"
const collectionName = "users"

var Collection *mongo.Collection

func New() {
	mongoURI := utils.GetEnvVars("MONGO_URI")

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Connected to MONGO_DB")
}
