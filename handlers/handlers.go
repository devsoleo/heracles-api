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
	if category == "" || query == "" || locale == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'category', 'query' and 'lang' are required"})
		return
	}

	creatures, err := services.Search(category, query, locale, c.Query("parentId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, creatures)
}

func HandleGenerate(c *gin.Context) {

}
