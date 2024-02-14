package repository

import (
	"context"
	"crud/internal/models"
	"crud/internal/postgres/usecase"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	postgresConnection *pgx.Conn
}

func NewRepository(conn *pgx.Conn) usecase.RepositoryInterfaces {
	return &repository{postgresConnection: conn}
}

func (r *repository) CreatePostgres(name string, length int, author string, album string) error {

	answer, err := r.postgresConnection.Exec(context.Background(), fmt.Sprintf("insert into songs(Name, Length, Author, Album) values ('%s', %d, '%s', '%s');", name, length, author, album))
	if err != nil {
		log.Println(err)
	}

	log.Println("Answer from insert", answer)

	return nil
}

func (r *repository) ReadPostgres() ([]models.PostgresData, error) {
	count := r.getCount()

	answer, err := r.postgresConnection.Query(context.Background(), "select Name, Length, Author, Album from songs;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer answer.Close()

	resultArr := make([]models.PostgresData, 0, count)

	for answer.Next() {
		var name, author, album, strLen string
		var length int

		answer.Scan(&name, &strLen, &author, &album)
		length, _ = strconv.Atoi(strLen)

		resultArr = append(resultArr, models.PostgresData{SongName: name, SongLength: length, SongAuthor: author, SongAlbum: album})
	}

	return resultArr, nil
}

func (r *repository) getCount() int {
	countRow, err := r.postgresConnection.Query(context.Background(), "select count(*) from songs;")

	if err != nil {
		log.Println(err)
		return -1
	}
	defer countRow.Close()

	countRow.Next()
	var count int
	countRow.Scan(&count)

	return count

}

func (r *repository) UpdatePostgres(name, author, album string) error {
	answer, err := r.postgresConnection.Exec(context.Background(), fmt.Sprintf("update songs set Author = '%s', Album = '%s' where Name = '%s';", author, album, name))
	if err != nil {
		log.Println(err)
	}

	log.Println("Answer from update", answer)

	return nil
}

func (r *repository) DeletePostgres(name, author, album string) error {
	answer, err := r.postgresConnection.Exec(context.Background(), fmt.Sprintf("delete from songs where Name = '%s' and Author = '%s' and Album = '%s';", name, author, album))

	if err != nil {
		return err
	}

	log.Println(answer)

	return nil
}
