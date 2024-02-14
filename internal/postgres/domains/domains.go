package domains

import (
	"crud/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	CreatePostgres(models.PostgresData) error
	ReadPostgres() ([]models.PostgresData, error)
	UpdatePostgres(models.PostgresData) error
	DeletePostgres(models.PostgresData) error
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

// ShowAccount godoc
//
//	@Summary		Create postgres data
//	@Tags			Postgres
//	@Description	You can add data to postgres database via this endpoint
//	@ID				create-postgres
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	PostgresDTO	true	"Json for creating data"
//	@Router			/createpostgres [post]
func (h *handler) CreatePostgres(ctx echo.Context) error {

	var data PostgresDTO

	if err := ctx.Bind(&data); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	modelData := &models.PostgresData{SongName: data.Name, SongLength: data.Length, SongAuthor: data.Author, SongAlbum: data.Album}

	err := h.usecase.CreatePostgres(*modelData)

	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Create postgres ok"})

}

// ShowAccount godoc
//
//	@Summary		Read postgres data
//	@Tags			Postgres
//	@Description	You can read data from postgres database via this endpoint
//	@ID				read-postgres
//	@Produce		json
//	@Router			/readpostgres [get]
func (h *handler) ReadPostgres(ctx echo.Context) error {
	rawData, err := h.usecase.ReadPostgres()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: rawData})
}

// ShowAccount godoc
//
//	@Summary		Update postgres data
//	@Tags			Postgres
//	@Description	You can update postgres data via this endpoint
//	@ID				update-postgres
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	PostgresDTO	true	"Json for updating data"
//	@Router			/updatepostgres [put]
func (h *handler) UpdatePostgres(ctx echo.Context) error {
	var jsonData PostgresDTO

	if err := ctx.Bind(&jsonData); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "JSON error"})
	}

	data := &models.PostgresData{SongName: jsonData.Name, SongLength: jsonData.Length, SongAuthor: jsonData.Author, SongAlbum: jsonData.Album}

	if err := h.usecase.UpdatePostgres(*data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Update ok"})
}

// ShowAccount godoc
//
//	@Summary		Delete postgres data
//	@Tags			Postgres
//	@Description	You can delete data from postgres database via this endpoint
//	@ID				delete-postgres
//	@Accept			json
//	@Produce		json
//	@Param			JsonExample	body	PostgresDTO	true	"Json for deleting data"
//	@Router			/deletepostgres [delete]
func (h *handler) DeletePostgres(ctx echo.Context) error {
	var jsonData PostgresDTO

	if err := ctx.Bind(&jsonData); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "JSON error"})
	}

	data := &models.PostgresData{SongName: jsonData.Name, SongLength: jsonData.Length, SongAuthor: jsonData.Author, SongAlbum: jsonData.Album}

	if err := h.usecase.DeletePostgres(*data); err != nil {
		return ctx.JSON(500, &models.Response{Status: 500, Payload: "Server error"})
	}

	return ctx.JSON(200, &models.Response{Status: 200, Payload: "Delete ok"})
}

func (h *handler) CreateRoutes(e *echo.Echo) {
	e.POST("/createpostgres", h.CreatePostgres)
	e.GET("/readpostgres", h.ReadPostgres)
	e.PUT("/updatepostgres", h.UpdatePostgres)
	e.DELETE("/deletepostgres", h.DeletePostgres)
}
