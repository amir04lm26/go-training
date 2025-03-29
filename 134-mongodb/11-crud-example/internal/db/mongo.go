package db

import (
	"context"
	"crud-example/config"
	ctxhelper "crud-example/pkg/util/context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Singleton instance
var (
	MongoClient *mongo.Client
	once        sync.Once
)

func InitMongoClient() {
	once.Do(func() {
		// "mongodb://bk-admin:123456@localhost:27017/bookstore"
		conStr := fmt.Sprintf("mongodb://%s:%s@%s/bookstore",
			config.GetEnv(config.DBUser),
			config.GetEnv(config.DBPassword),
			config.GetEnv(config.DBHost))
		ctx, cancel := ctxhelper.CtxWithNormalTimeout()
		defer cancel()

		// clientOptions := options.Client().ApplyURI(conStr)
		// Set connection pool options
		clientOptions := options.Client().
			ApplyURI(conStr).
			SetMaxPoolSize(7).                    // Limit to 7 connections
			SetMinPoolSize(2).                    // Keep at least 2 connections open
			SetMaxConnIdleTime(30 * time.Second). // Close idle connections after 30s
			SetMaxConnecting(5)                   // Allow 5 concurrent new connections

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Panicf("MongoDB connection failed: %v\n", err)
		}

		if err := client.Ping(ctx, nil); err != nil {
			log.Panicf("MongoDB ping failed: %v\n", err)
		}

		MongoClient = client
		fmt.Println("Connected to MongoDB successfully!")
	})
}

func DisconnectMongoClient() {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error disconnecting MongoDB:", err)
	}
	fmt.Println("Disconnected from MongoDB.")
}
