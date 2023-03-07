package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"url-api/models"
)

func GetMaxChanPrice(priceChan <-chan models.Price) models.Price {

	var max models.Price
	for price := range priceChan {
		if price.Price > max.Price {
			max = price
		}
	}
	return max
}

func GetIntSlice(arr []string) []int {
	var intArr = []int{}

	for _, i := range arr {
		j, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}

		intArr = append(intArr, j)
	}

	return intArr
}

func GetUrlById(items []models.Url, id int) (models.Url, bool) {
	for _, item := range items {
		if item.Id == id {
			return item, true
		}
	}
	return models.Url{}, false
}

func HandlePanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("")
	}
}
