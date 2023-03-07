package repositories

import (
	"url-api/models"

	"github.com/jinzhu/gorm"
)

type UrlsRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlsRepository {
	return &UrlsRepository{db: db}
}

func (r *UrlsRepository) Close() {
	r.db.Close()
}

func (r *UrlsRepository) GetUrls() []models.Url {

	var urls []models.Url
	r.db.Find(&urls)

	return urls
}
