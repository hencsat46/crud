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
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) ReadMongo(ctx echo.Context) error {
	result, err := h.usecase.ReadMongo()

	if err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: result})
}

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

func (h *handler) CreateRoutes(e *echo.Echo) {
	e.GET("/readmongo", h.ReadMongo)
	e.POST("createmongo", h.CreateMongo)
}
