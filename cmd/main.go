package main

import (
	"crud/internal/initdb"
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

	connection, err := initdb.InitPostgres()
	redisConnection := initdb.InitRedis()

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

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	handler.CreateRoutes(e)
	domainsRedis.CreateRoutes(e)

	e.Start(":3000")
}
