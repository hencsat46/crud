package initdb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const postgresUrl = "postgres://postgres:forstudy@localhost:5432/temp"
const mongodbUrl = "mongodb://localhost:27017"

func InitPostgres() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), postgresUrl)

	if err != nil {
		log.Fatal("Cannot connect to postgres")
	}

	return conn, nil

}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func InitMongo() *mongo.Client {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodbUrl).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("Cannot connect to mongoDB")
	}

	return client

}
