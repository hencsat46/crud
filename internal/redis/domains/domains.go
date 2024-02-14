package domains

import (
	"crud/internal/models"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type domains struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	CreateRedis(models.RedisData) error
	ReadRedis(models.RedisData) (string, error)
	UpdateRedis(models.RedisData) error
	DeleteRedis(models.RedisData) error
}

func NewDomains(usecase UsecaseInterfaces) *domains {
	return &domains{usecase: usecase}
}

// ShowAccount godoc
//
//	@Summary		Create redis data
//	@Tags			Redis
//	@Description	You can add data to redis database via this endpoint
//	@ID				create-redis
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	redisDTO	true	"Json for creating data"
//	@Router			/createredis [post]
func (d *domains) CreateRedis(ctx echo.Context) error {

	var jsonData redisDTO

	if err := ctx.Bind(&jsonData); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	data := models.RedisData{Key: jsonData.Key, Value: jsonData.Value}

	if err := d.usecase.CreateRedis(data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Create redis ok"})
}

// ShowAccount godoc
//
//	@Summary		Read redis data
//	@Tags			Redis
//	@Description	You can read data from redis database via this endpoint
//	@ID				read-redis
//	@Accept			json
//	@Produce		json
//	@Param			key	query	string	true	"name search by query"
//	@Router			/readredis [get]
func (d *domains) ReadRedis(ctx echo.Context) error {
	// var jsonData redisDTO

	// if err := ctx.Bind(&jsonData); err != nil {
	// 	log.Println(err)
	// 	return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	// }

	key := ctx.QueryParam("key")
	if key == "" {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad query"})
	}

	data := models.RedisData{Key: key}

	result, err := d.usecase.ReadRedis(data)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ctx.JSON(404, &models.Response{Status: 404, Payload: "Not found"})
		}
		log.Println(err)
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: result})

}

// ShowAccount godoc
//
//	@Summary		Update redis data
//	@Tags			Redis
//	@Description	You can update redis data via this endpoint
//	@ID				update-redis
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	redisDTO	true	"Json for updating data"
//	@Router			/updateredis [put]
func (d *domains) UpdateRedis(ctx echo.Context) error {

	var jsonData redisDTO

	if err := ctx.Bind(&jsonData); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	data := models.RedisData{Key: jsonData.Key, Value: jsonData.Value}

	if err := d.usecase.UpdateRedis(data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}
	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Update ok"})
}

// ShowAccount godoc
//
//	@Summary		Delete redis data
//	@Tags			Redis
//	@Description	You can delete data from redis database via this endpoint
//	@ID				delete-redis
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	redisDTO	true	"Json for deleting data"
//	@Router			/deleteredis [delete]
func (d *domains) DeleteRedis(ctx echo.Context) error {

	var jsonData redisDTO

	if err := ctx.Bind(&jsonData); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	data := models.RedisData{Key: jsonData.Key, Value: jsonData.Value}

	if err := d.usecase.DeleteRedis(data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}
	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Delete ok"})
}

func (d *domains) CreateRoutes(e *echo.Echo) {
	e.POST("/createredis", d.CreateRedis)
	e.GET("/readredis", d.ReadRedis)
	e.PUT("/updateredis", d.UpdateRedis)
	e.DELETE("/deleteredis", d.DeleteRedis)
}
