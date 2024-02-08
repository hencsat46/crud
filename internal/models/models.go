package models

type PostgresData struct {
	SongName   string
	SongLength int
	SongAuthor string
	SongAlbum  string
}

type RedisData struct {
	Key   string
	Value string
}

type Response struct {
	Status  int
	Payload interface{}
}
