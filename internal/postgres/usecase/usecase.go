package usecase

import (
	"crud/internal/models"
	"crud/internal/postgres/domains"
	"log"
)

type usecase struct {
	repository RepositoryInterfaces
}

type RepositoryInterfaces interface {
	CreatePostgres(string, int, string, string) error
	ReadPostgres() ([]models.PostgresData, error)
	UpdatePostgres(string, string, string) error
	DeletePostgres(string, string, string) error
}

func NewUsecase(repo RepositoryInterfaces) domains.UsecaseInterfaces {
	return &usecase{repository: repo}
}

func (u *usecase) CreatePostgres(data models.PostgresData) error {
	err := u.repository.CreatePostgres(data.SongName, data.SongLength, data.SongAuthor, data.SongAlbum)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) ReadPostgres() ([]models.PostgresData, error) {
	result, err := u.repository.ReadPostgres()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *usecase) UpdatePostgres(data models.PostgresData) error {
	if err := u.repository.UpdatePostgres(data.SongName, data.SongAuthor, data.SongAlbum); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) DeletePostgres(data models.PostgresData) error {
	if err := u.repository.DeletePostgres(data.SongName, data.SongAuthor, data.SongAlbum); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
