package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Singleton instance
var (
	clientInstance *mongo.Client
	once           sync.Once
)

func GetMongoClient() *mongo.Client {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// clientOptions := options.Client().ApplyURI("mongodb://root:123456@localhost:27017")
		// Set connection pool options
		clientOptions := options.Client().
			ApplyURI("mongodb://root:123456@localhost:27017").
			SetMaxPoolSize(50).                   // Limit to 50 connections
			SetMinPoolSize(10).                   // Keep at least 10 connections open
			SetMaxConnIdleTime(30 * time.Second). // Close idle connections after 30s
			SetMaxConnecting(5)                   // Allow 5 concurrent new connections

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Panicf("MongoDB connection failed: %v\n", err)
		}

		if err := client.Ping(ctx, nil); err != nil {
			log.Panicf("MongoDB ping failed: %v\n", err)
		}

		clientInstance = client
		fmt.Println("Connected to MongoDB successfully!")
	})

	return clientInstance
}

func DisconnectMongoClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error disconnecting MongoDB:", err)
	}
	fmt.Println("Disconnected from MongoDB.")
}
