package repository

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func ConnectDB() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster.mnyoqzm.mongodb.net/?retryWrites=true&w=majority&appName=cluster",
		os.Getenv("USER_NAME"),
		os.Getenv("PASSWORD"),
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return client, nil
}
