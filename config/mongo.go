package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient          *mongo.Client
	RumahSakitCollection *mongo.Collection
)

func ConnectMongo() {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ Gagal konek MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB tidak merespon:", err)
	}

	MongoClient = client
	RumahSakitCollection = client.Database(dbName).Collection("rumah_sakit")

	log.Println("✅ MongoDB connected")
}
