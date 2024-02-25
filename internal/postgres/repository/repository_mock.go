package repository

import (
	"crud/internal/models"
	"crud/internal/postgres/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type tempPostgres struct {
	Name   string
	Length int
	Author string
	Album  string
}

type repositoryMock struct {
	data []tempPostgres
}

func NewMock() usecase.RepositoryInterfaces {
	return &repositoryMock{data: make([]tempPostgres, 0, 20)}
}

func (r *repositoryMock) CreatePostgres(name string, length int, author, album string) error {
	r.data = append(r.data, tempPostgres{Name: name, Length: length, Author: author, Album: album})
	return nil
}

func (r *repositoryMock) ReadPostgres() ([]models.PostgresData, error) {
	result := make([]models.PostgresData, 20)

	for _, value := range r.data {
		result = append(result, models.PostgresData{SongName: value.Name, SongLength: value.Length, SongAuthor: value.Author, SongAlbum: value.Album})
	}

	return result, nil
}

func (r *repositoryMock) UpdatePostgres(name, author, album string) error {
	index, ok := r.checkElem(name, author, album)

	if ok {
		r.data[index] = tempPostgres{Name: name, Author: author, Album: album}
		return nil
	}

	return status.Error(codes.NotFound, "not found")
}

func (r *repositoryMock) checkElem(name, author, album string) (int, bool) {
	for i := 0; i < len(r.data); i++ {
		if r.data[i].Name == name && r.data[i].Author == author && r.data[i].Album == album {
			return i, true
		}
	}

	return -1, false
}

func (r *repositoryMock) DeletePostgres(name, author, album string) error {
	index, ok := r.checkElem(name, author, album)

	if ok {
		remove(r.data, index)
		return nil
	}

	return status.Error(codes.NotFound, "not found")
}

func remove(slice []tempPostgres, s int) []tempPostgres {
	return append(slice[:s], slice[s+1:]...)
}
