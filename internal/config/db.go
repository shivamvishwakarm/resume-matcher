package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Exported package-level variables — accessible as config.Db, config.Client, etc.
var (
	Db       *mongo.Database
	Client   *mongo.Client
	MongoCtx context.Context
	Cancel   context.CancelFunc
)

const MONGO_URI = "mongodb://admin:password@localhost:27017/?authSource=admin"
const DbName = "resume-matcher"
const ColName = "user"

func init() {

	client, err := mongo.Connect(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping to verify the connection is alive
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	Client = client
	Db = client.Database(DbName)
	MongoCtx = context.Background()

	log.Println("Connected to MongoDB successfully")
}
