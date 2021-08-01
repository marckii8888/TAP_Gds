package web

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/marckii8888/TAP_Gds/Backend/config"
	"github.com/marckii8888/TAP_Gds/Backend/web/handlers"
	"log"
	"strconv"
)

type Router struct {
	*gin.Engine
}

var PORT uint64

func NewRouter() *Router {
	router := gin.Default()

	// Enable CORS as middleware
	router.Use(cors.Default())

	// Read config file
	config.InitConf()
	cfgPort, err := strconv.ParseUint(config.Conf.Server.Port, 10, 64)
	if err != nil {
		log.Fatal("Failed to convert from config file")
	}
	PORT = cfgPort

	helper := handlers.New()
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
	err := r.Engine.Run(fmt.Sprintf(":%+v", PORT))
	if err != nil {
		log.Fatal("Failed to start router")
	}
	log.Printf("Connected to port %+v", PORT)
}

func Run(){
	NewRouter().Run()
}