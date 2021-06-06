package handler

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response := "pong"
	GenerateResponse(c, response, nil)
}

//func Test(c *gin.Context) {
//	_, err := utils.GenerateMySQLDSN()
//	GenerateEmptyDataResponse(c, err)
//}
