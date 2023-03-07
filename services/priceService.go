package services

import (
	"fmt"
	"sync"
	"url-api/models"
	"url-api/repositories"
	"url-api/utils"
)

type PriceService struct {
	Repo repositories.UrlsRepository
}

func (serv *PriceService) GetMaxPrice(urlPackage []int) (models.Price, error) {

	var wg sync.WaitGroup

	urls := serv.Repo.GetUrls()
	priceChan := make(chan models.Price, len(urlPackage))

	for _, urlId := range urlPackage {

		wg.Add(1)

		priceUrl, exists := utils.GetUrlById(urls, urlId)

		if !exists {
			return models.Price{}, fmt.Errorf("url not found")
		}

		go func(url string) {

			defer wg.Done()
			priceChan <- utils.SendGetRequest[models.Price](url)

		}(priceUrl.Url)
	}

	wg.Wait()
	close(priceChan)

	return utils.GetMaxChanPrice(priceChan), nil
}
