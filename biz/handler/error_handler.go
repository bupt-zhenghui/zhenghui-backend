package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "success",
			"data":    data,
		})
	}
}

func GenerateEmptyDataResponse(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "success",
		})
	}
}
