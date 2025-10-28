package config

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://18223038_db_user:33wxubHLymQ7kRmD@pawm.2adoicu.mongodb.net/").SetServerSelectionTimeout(30 * time.Second)

	// Optional: enable insecure TLS skip for debugging (only when explicitly set).
	// To enable set environment variable MONGO_INSECURE_TLS=true or 1.
	if v := os.Getenv("MONGO_INSECURE_TLS"); v != "" {
		if strings.ToLower(v) == "1" || strings.ToLower(v) == "true" {
			tlsConfig := &tls.Config{InsecureSkipVerify: true}
			clientOptions.SetTLSConfig(tlsConfig)
			log.Println("WARNING: MongoDB TLS verification disabled (MONGO_INSECURE_TLS=true). Use only for debugging.")
		}
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("user")
	log.Println("Connected to MongoDB!")
}
