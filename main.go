package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"url-api/db"
	"url-api/models"
	"url-api/repositories"
	"url-api/services"
	"url-api/utils"

	"github.com/gorilla/mux"
)

var priceService services.PriceService

const port = ":8080"

func main() {

	defer priceService.Repo.Close()

	priceService = services.PriceService{
		Repo: *repositories.NewUrlRepository(db.OpenDb()),
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/price", getPrice).Methods("GET")
	router.HandleFunc("/price", postPrice).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))
}

func getPrice(w http.ResponseWriter, r *http.Request) {

	defer utils.HandlePanic(w)

	query := r.URL.Query()

	requestId, _ := strconv.Atoi(query.Get("request_id"))
	urlPackage := utils.GetIntSlice(query["url_package"])
	ip := query.Get("ip")

	handleRequest(w, requestId, urlPackage, ip)
}

func postPrice(w http.ResponseWriter, r *http.Request) {

	defer utils.HandlePanic(w)

	var body models.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		panic("No content")
	}

	handleRequest(w, body.RequestId, body.UrlPackage, body.IP)
}

func handleRequest(w http.ResponseWriter, requestId int, urlPackage []int, ip string) {

	defer utils.HandlePanic(w)

	if !isUrlRequestValid(urlPackage, ip, requestId) {
		panic("No content")
	}

	maxPrice, err := priceService.GetMaxPrice(urlPackage)

	if err != nil {
		panic("No content")
	}

	json.NewEncoder(w).Encode(maxPrice)
}

func isUrlRequestValid(urlPackage []int, ip string, requestId int) bool {
	return len(urlPackage) != 0 && net.ParseIP(ip) != nil
}
