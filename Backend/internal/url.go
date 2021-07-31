package internal

import (
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type URL struct{
	OriginalUrl string `json:"original_url"`
	ShortUrl string `json:"short_url"`
	Code string `json:"code"`
}

type URLReq struct {
	OriginalUrl string `json:"original_url"`
}

const BASE_URL = "http://localhost:8081/"

// Add to DB
func ShortenURL(db *gorm.DB, originalUrl string) (string ,error) {
	// Generate a unique code
	uniqueCode := ksuid.New().String()[0:7]
	// Append unique code to localhost:8081/code
	shortUrl := BASE_URL + uniqueCode

	newURL := &URL{
		OriginalUrl: originalUrl,
		ShortUrl: shortUrl,
		Code: uniqueCode,
	}

	err := db.Create(newURL).Error
	if err != nil {
		return "",err
	}
	return shortUrl, nil
}

// Retrieve from DB
func RedirectURL(db *gorm.DB, code string, url *URL) error {
	err := db.Where("code = ?", code).First(url).Error
	if err != nil {
		return err
	}
	return nil
}