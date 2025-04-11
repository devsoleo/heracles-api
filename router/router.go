package router

import (
	"devsoleo/heracles-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/search", handlers.HandleSearch)
		api.GET("/generate", handlers.HandleGenerate)
	}

	return r
}
