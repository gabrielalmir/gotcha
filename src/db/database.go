package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	uri    string
}

func NewDatabase(uri string) *Database {
	return &Database{
		uri: uri,
	}
}

func (d *Database) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(d.uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	d.client = client
	log.Println("Connected to MongoDB!")
	return nil
}

func (d *Database) GetCollection(collectionName string) *mongo.Collection {
	return d.client.Database("gotcha").Collection(collectionName)
}
