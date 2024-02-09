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

func (h *handler) CreateRoutes(e *echo.Echo) {
	e.GET("/readmongo", h.ReadMongo)
}
