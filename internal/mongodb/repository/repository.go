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

	filter := bson.M{}

	cursor, err := r.mongoConnection.Database("temp").Collection("temp").Find(context.Background(), filter, options.Find())
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

	for cursor.Next(context.Background()) {
		var elem models.MongoData
		if err := cursor.Decode(&elem); err != nil {
			log.Println("Парсингу курсора пизда")
			return nil, err
		}

		results = append(results, elem)
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

func (r *repository) UpdateMongo(data models.MongoData) error {

	filter := bson.D{{"key", data.Key}}

	update := bson.D{{"$set", bson.D{{"value", data.Value}}}}

	if _, err := r.mongoConnection.Database("temp").Collection("temp").UpdateOne(context.Background(), filter, update); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *repository) DeleteMongo(data models.MongoData) error {
	filter := bson.D{{"key", data.Key}}

	if _, err := r.mongoConnection.Database("temp").Collection("temp").DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}
