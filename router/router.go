package router

import (
    "github.com/gin-gonic/gin"
    "devsoleo/heracles-back/handlers"
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
