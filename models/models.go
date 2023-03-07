package models

type Price struct {
	Price float32 `json:"price"`
}

type RequestBody struct {
	RequestId  int    `json:"request_id"`
	UrlPackage []int  `json:"url_package"`
	IP         string `json:"ip"`
}

type Url struct {
	Id  int    `gorm:"primaryKey"`
	Url string `gorm:"not null"`
}
