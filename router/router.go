package router

import (
	"devsoleo/heracles-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type", "Access-Control-Allow-Origin"},
	}))

	api := r.Group("/api")
	{
		api.GET("/search", handlers.HandleSearch)
		api.GET("/generate", handlers.HandleGenerate)
	}

	return r
}
