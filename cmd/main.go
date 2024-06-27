package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zeze322/wt-guided-weaponry/api"
	"github.com/zeze322/wt-guided-weaponry/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		port              = os.Getenv("PORT")
		mysqlURL          = os.Getenv("MYSQL_URL")
		mongodbURL        = os.Getenv("MONGODB_URL")
		mongodbDatabase   = os.Getenv("MONGODB_DATABASE")
		mongodbCollection = os.Getenv("MONGODB_COLLECTION")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mysqlStore, err := storage.NewMySQLStorage(mysqlURL)
	if err != nil {
		log.Println(err)
	}
	mongodbStore, err := storage.NewMongoDBStorage(ctx, mongodbURL, mongodbDatabase, mongodbCollection)
	if err != nil {
		log.Println(err)
	}

	s := api.NewServer(port, mysqlStore, mongodbStore)
	s.Start()
}
