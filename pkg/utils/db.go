package utils

import (
	"context"
	"fmt"
	"log"

	"aas.dev/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatabase *mongo.Database

type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func init() {
	cfg := config.LoadConfig()
	ConnectDB(cfg)
}

func ConnectDB(cfg *config.Config) {

	// Set the Stable API version for MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.DBURI).SetServerAPIOptions(serverAPI)

	// Create a new MongoDB client and connect
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to confirm the connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// If successful, log and initialize the MongoDatabase variable
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	// MongoDatabase = client.Database(cfg.DBName) // Use the DB name from config
	MongoDatabase = client.Database("myapp") // Use the DB name from config

}
