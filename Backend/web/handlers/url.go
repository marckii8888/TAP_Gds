package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	internal2 "github.com/marckii8888/TAP_Gds/Backend/internal"
	"gorm.io/gorm"
	"net/http"
)

type Helper struct{
	db *gorm.DB
}

func New() *Helper {
	db := internal2.InitDB()
	db.AutoMigrate(&internal2.URL{})
	return &Helper{db : db}
}

func (helper *Helper) ShortenURL(c *gin.Context){
	var req internal2.URLReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : fmt.Sprintf("Error : %+v", err),
		})
		return
	}
	fmt.Printf("%+v", req)
	result, err := internal2.ShortenURL(helper.db, req.OriginalUrl)
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

func (helper *Helper) Redirect(c *gin.Context){
	code, _ := c.Params.Get("code")
	var url internal2.URL
	err := internal2.RedirectURL(helper.db, code, &url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : fmt.Sprintf("Error : %+v", err),
		})
		return
	}
	c.Redirect(http.StatusFound, url.OriginalUrl)
}
