package usecase

import (
	"crud/internal/models"
	"crud/internal/redis/domains"
	"log"
)

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	CreateRedis(string, string) error
	ReadRedis(string) (string, error)
	UpdateRedis(string, string) error
	DeleteRedis(string) error
}

func NewUsecase(repo RepositoryInterfaces) domains.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) CreateRedis(data models.RedisData) error {
	if err := u.repo.CreateRedis(data.Key, data.Value); err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (u *usecase) ReadRedis(data models.RedisData) (string, error) {
	result, err := u.repo.ReadRedis(data.Key)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return result, nil
}

func (u *usecase) UpdateRedis(data models.RedisData) error {
	if err := u.repo.UpdateRedis(data.Key, data.Value); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) DeleteRedis(data models.RedisData) error {
	if err := u.repo.DeleteRedis(data.Key); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
