package usecase

import (
	"crud/internal/models"
	"crud/internal/mongodb/handler"
	"log"
)

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	CreateMongo(models.MongoData) error
	ReadMongo() ([]models.MongoData, error)
}

func NewUsecase(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) ReadMongo() ([]models.MongoData, error) {
	result, err := u.repo.ReadMongo()

	if err != nil {
		log.Println(err)
	}

	return result, nil
}

func (u *usecase) CreateMongo(data models.MongoData) error {
	if err := u.repo.CreateMongo(data); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
