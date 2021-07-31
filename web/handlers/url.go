package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/marckii8888/TAP_Gds/internal"
	"gorm.io/gorm"
	"net/http"
)

type Helper struct{
	db *gorm.DB
}

func New() *Helper{
	db := internal.InitDB()
	db.AutoMigrate(&internal.URL{})
	return &Helper{db : db}
}

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
			"message" : fmt.Sprintf("Error : %+v"),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : result,
	})
}

func (helper *Helper) Redirect(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Redirected",
	})
}
