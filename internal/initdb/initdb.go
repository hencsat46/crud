package initdb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

const postgresUrl = "postgres://postgres:forstudy@localhost:5432/temp"

func InitPostgres() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), postgresUrl)

	if err != nil {
		log.Println(err)
		return nil, err
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
