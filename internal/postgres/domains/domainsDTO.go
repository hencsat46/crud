package domains

type PostgresDTO struct {
	Name   string `json:"Name"`
	Length int    `json:"Length"`
	Author string `json:"Author"`
	Album  string `json:"Album"`
}
