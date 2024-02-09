package repository

import (
	"context"
	"crud/internal/models"
	"crud/internal/mongodb/usecase"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	mongoConnection *mongo.Client
}

func NewRepository(conn *mongo.Client) usecase.RepositoryInterfaces {
	return &repository{mongoConnection: conn}
}

func (r *repository) CreateMongo(value models.MongoData) error {
	if _, err := r.mongoConnection.Database("temp").Collection("temp").InsertOne(context.Background(), value); err != nil {
		return err
	}

	return nil
}

func (r *repository) ReadMongo() ([]models.MongoData, error) {
	cursor, err := r.mongoConnection.Database("temp").Collection("temp").Find(context.Background(), options.Find().SetSort(bson.D{}))
	if err != nil {
		log.Println("haha")
		return nil, err
	}

	count, err := r.countDocuments()
	if err != nil {
		log.Println("hui")
		return nil, err
	}
	results := make([]models.MongoData, count)

	if err = cursor.All(context.Background(), &results); err != nil {
		log.Println("pizda")
		return nil, err
	}

	return results, nil
}

func (r *repository) countDocuments() (int64, error) {
	count, err := r.mongoConnection.Database("temp").Collection("temp").CountDocuments(context.Background(), options.Find().SetSort(bson.D{}))

	if err != nil {
		return -1, err
	}

	return count, nil
}
