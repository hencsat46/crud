package domains

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// func TestCreatePostgres(t *testing.T) {

// 	// testcases := []struct {
// 	// 	request struct

// 	// }

// }

func temp() {
	e := echo.New()

	tempBody := struct {
		Name   string `json:"Name"`
		Length int    `json:"Length"`
		Author string `json:"Author"`
		Album  string `json:"Album"`
	}{
		Name:   "lsdkfj",
		Length: 43,
		Author: "lsdkfj",
		Album:  "flskdjf",
	}

	j, err := json.Marshal(tempBody)

	if err != nil {
		log.Println(err)
		return
	}

	read := strings.NewReader()

	request, err := http.NewRequest("POST", "localhost:3000", tempBody)
}
