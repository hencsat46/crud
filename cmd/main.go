package main

import (
	"crud/internal/initdb"
	handlerMongo "crud/internal/mongodb/handler"
	repositoryMongo "crud/internal/mongodb/repository"
	usecaseMongo "crud/internal/mongodb/usecase"

	domainsPostgres "crud/internal/postgres/domains"
	repositoryPostgres "crud/internal/postgres/repository"
	usecasePostgres "crud/internal/postgres/usecase"

	domainsRedis "crud/internal/redis/domains"
	repositoryRedis "crud/internal/redis/repository"
	usecaseRedis "crud/internal/redis/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	initdb.InitMongo()

	connection, err := initdb.InitPostgres()
	redisConnection := initdb.InitRedis()
	mongoConnection := initdb.InitMongo()

	if err != nil {
		log.Println(err)
		return
	}

	repo := repositoryPostgres.NewRepository(connection)
	usecase := usecasePostgres.NewUsecase(repo)
	handler := domainsPostgres.NewHandler(usecase)

	repoRedis := repositoryRedis.NewRepository(redisConnection)
	usecaseRedis := usecaseRedis.NewUsecase(repoRedis)
	domainsRedis := domainsRedis.NewDomains(usecaseRedis)

	repoMongo := repositoryMongo.NewRepository(mongoConnection)
	usecaseMongo := usecaseMongo.NewUsecase(repoMongo)
	handlerMongo := handlerMongo.NewHandler(usecaseMongo)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	handler.CreateRoutes(e)
	domainsRedis.CreateRoutes(e)
	handlerMongo.CreateRoutes(e)

	e.Start(":3000")
}
