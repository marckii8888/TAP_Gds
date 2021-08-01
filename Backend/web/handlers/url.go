package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/marckii8888/TAP_Gds/Backend/internal"
	"gorm.io/gorm"
	"net/http"
)

type Helper struct{
	db *gorm.DB
}

func New() *Helper {
	db := internal.InitDB()
	db.AutoMigrate(&internal.URL{})
	return &Helper{db : db}
}

// ShortenURL
// @Summary handler to generate a unique code and store it in the db with the original url
func (helper *Helper) ShortenURL(c *gin.Context){
	var req internal.URLReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : fmt.Sprintf("Error : %+v", err),
		})
		return
	}
	result, err := internal.ShortenURL(helper.db, req.OriginalUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : fmt.Sprintf("Error : %+v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : result,
	})
}

// Redirect
// @Summary helper function that processes a GET request. Looks up
// the unique code in the database and retrieve the original url
// @Params unique code
func (helper *Helper) Redirect(c *gin.Context){
	code, _ := c.Params.Get("code")
	var url internal.URL
	err := internal.RedirectURL(helper.db, code, &url)
	if err != nil {
		// If the shortened url does not exists
		c.JSON(http.StatusOK, gin.H{
			"message" : "Error! Invalid url",
		})
		return
	}
	c.Redirect(http.StatusFound, url.OriginalUrl)
}
