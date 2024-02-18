package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRequest[T any](c *gin.Context) (T, bool) {
	var request T

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return request, false
	}

	return request, true
}
