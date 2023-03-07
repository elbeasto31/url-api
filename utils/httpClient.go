package utils

import (
	"encoding/json"
	"net/http"
)

func SendGetRequest[T comparable](url string) T {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var item T
	json.NewDecoder(response.Body).Decode(&item)

	return item
}
