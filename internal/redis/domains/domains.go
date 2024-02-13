package domains

import (
	"crud/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
//	@Description	You can add data to redis database via this endpoint
//	@ID				get-string-by-int
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
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

func (d *domains) ReadRedis(ctx echo.Context) error {
	var jsonData redisDTO

	if err := ctx.Bind(&jsonData); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	data := models.RedisData{Key: jsonData.Key, Value: jsonData.Value}

	result, err := d.usecase.ReadRedis(data)
	if err != nil {
		log.Println(err)
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: result})

}

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
