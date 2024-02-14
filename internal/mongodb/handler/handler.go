package handler

import (
	"crud/internal/models"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	ReadMongo() ([]models.MongoData, error)
	CreateMongo(models.MongoData) error
	UpdateMongo(models.MongoData) error
	DeleteMongo(models.MongoData) error
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

// ShowAccount godoc
//
//	@Summary		Read mongo data
//	@Tags			Mongo
//	@Description	You can read data from mongo database via this endpoint
//	@ID				read-mongo
//	@Produce		json
//	@Router			/readmongo [get]
func (h *handler) ReadMongo(ctx echo.Context) error {
	result, err := h.usecase.ReadMongo()

	if err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: result})
}

// ShowAccount godoc
//
//	@Summary		Create mongo data
//	@Tags			Mongo
//	@Description	You can add data to mongo database via this endpoint
//	@ID				create-mongo
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	MongoDTO	true	"Json for creating data"
//	@Router			/createmongo [post]
func (h *handler) CreateMongo(ctx echo.Context) error {
	var rawData MongoDTO

	if err := ctx.Bind(&rawData); err != nil {
		return ctx.JSON(400, &models.Response{Status: 400, Payload: "Bad json"})
	}

	data := models.MongoData{Key: rawData.Key, Value: rawData.Value}

	if err := h.usecase.CreateMongo(data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Create ok"})
}

// ShowAccount godoc
//
//	@Summary		Update mongo data
//	@Tags			Mongo
//	@Description	You can update mongo data via this endpoint
//	@ID				update-mongo
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	MongoDTO	true	"Json for updating data"
//	@Router			/updatemongo [put]
func (h *handler) UpdateMongo(ctx echo.Context) error {
	var rawData MongoDTO

	if err := ctx.Bind(&rawData); err != nil {
		return ctx.JSON(400, &models.Response{Status: 400, Payload: "Bad json"})
	}

	data := models.MongoData{Key: rawData.Key, Value: rawData.Value}

	if err := h.usecase.UpdateMongo(data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Update ok"})
}

// ShowAccount godoc
//
//	@Summary		Delete mongo data
//	@Tags			Mongo
//	@Description	You can delete data from mongo database via this endpoint
//	@ID				delete-mongo
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	MongoDTO	true	"Json for deleting data"
//	@Router			/deletemongo [delete]
func (h *handler) DeleteMongo(ctx echo.Context) error {
	var rawData MongoDTO

	if err := ctx.Bind(&rawData); err != nil {
		return ctx.JSON(400, &models.Response{Status: 400, Payload: "Bad json"})
	}

	data := models.MongoData{Key: rawData.Key, Value: rawData.Value}

	if err := h.usecase.DeleteMongo(data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Delete Ok"})
}

func (h *handler) CreateRoutes(e *echo.Echo) {
	e.GET("/readmongo", h.ReadMongo)
	e.POST("createmongo", h.CreateMongo)
	e.PUT("/updatemongo", h.UpdateMongo)
	e.DELETE("/deletemongo", h.DeleteMongo)
}
