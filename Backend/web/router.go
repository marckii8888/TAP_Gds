package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handlers2 "github.com/marckii8888/TAP_Gds/Backend/web/handlers"
	"log"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()

	helper := handlers2.New()
	urlAPI := router.Group("/api")
	// POST Request to return the shorten url
	urlAPI.POST("/shorten", helper.ShortenURL)
	// GET request to redirect
	router.GET("/:code", helper.Redirect)

	return &Router{
		router,
	}
}

func (r *Router) Run(){
	port := 8081
	err := r.Engine.Run(fmt.Sprintf(":%+v", port))
	if err != nil {
		log.Fatal("Failed to start router")
	}
	log.Printf("Connected to port %+v", port)
}

func Run(){
	NewRouter().Run()
}