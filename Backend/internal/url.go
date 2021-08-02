package internal

import (
	"errors"
	"fmt"
	"github.com/marckii8888/TAP_Gds/Backend/config"
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


// ShortenURL
// @Summary Function to shorten the url by generating unique code and storing in db
func ShortenURL(db *gorm.DB, originalUrl string) (string ,error) {
	var dbUrl URL
	var uniqueCode string
	for {
		// Generate the unique code
		uniqueCode = ksuid.New().String()[0:7]

		err := db.Where("code = ?", uniqueCode).First(&dbUrl).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If code does not exists in database
			break
		}
	}
	// Append unique code to localhost:8081/code
	var BASE_URL = fmt.Sprintf("http://%+v:+%v", config.Conf.Server.Host, config.Conf.Server.Port)
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

// RedirectURL
// @Summary Checks if unique code exists in database
func RedirectURL(db *gorm.DB, code string, url *URL) error {
	err := db.Where("code = ?", code).First(url).Error
	if err != nil {
		return err
	}
	return nil
}