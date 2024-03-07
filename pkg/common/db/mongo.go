package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoService struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoService() (*MongoService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	//2.选择数据库 my_db
	db := client.Database("apsc_im")
	if err != nil {
		return nil, err
	}
	return &MongoService{
		client,
		db,
	}, nil
}
