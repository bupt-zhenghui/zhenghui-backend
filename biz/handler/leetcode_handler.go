package handler

import (
	"github.com/gin-gonic/gin"
	"zhenghui-backend/biz/service"
)

func SearchLeetcodeHandler(c *gin.Context) {
	codeList, err := service.SearchLeetcode()
	GenerateResponse(c, codeList, err)
}

func SearchLeetcodeStatisticsHandler(c *gin.Context) {
	response, err := service.SearchLeetcodeStatistics()
	GenerateResponse(c, response, err)
}

func UpdateLeetcodeStatisticsHandler(c *gin.Context) {
	err := service.UpdateLeetcodeStatistics()
	GenerateEmptyDataResponse(c, err)
}
