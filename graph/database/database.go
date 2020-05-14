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

func New() *mongo.Database {
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

	db := client.Database(dbName)

	fmt.Println("Connected to MONGO_DB")

	return db
}
