package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"devsoleo/heracles-api/services"
)

func HandleSearch(c *gin.Context) {
	category := c.Query("category")
	query := c.Query("query")
	locale := c.Query("locale")

	// TODO : add parentid check
	if category == "" || locale == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'category' and 'locale' are required"})
		return
	}

	result, err := services.Search(category, query, locale, c.Query("parentId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func HandleGenerate(c *gin.Context) {
	var rawForms []map[string]interface{}

	if err := c.ShouldBindJSON(&rawForms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.Generate(rawForms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
