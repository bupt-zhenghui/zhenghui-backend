package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhenghui-backend/biz/consts"
	"zhenghui-backend/biz/service"
)

func Ping(c *gin.Context) {
	response := "pong"
	GenerateResponse(c, response, nil)
}

func Test(c *gin.Context) {
	GenerateEmptyDataResponse(c, nil)
}

func InsertAccessDataHandler(c *gin.Context) {
	clientIP := c.ClientIP()
	page, err := strconv.ParseInt(c.Param("page"), 10, 64)
	if err != nil {
		GenerateEmptyDataResponse(c, err)
	}
	accessPage := consts.AccessPage(page)
	err = service.InsertAccessData(clientIP, accessPage)
	GenerateEmptyDataResponse(c, err)
}

func SearchAccessHandler(c *gin.Context) {
	response, err := service.SearchAccessHistory()
	GenerateResponse(c, response, err)
}

func SearchAccessStatisticsHandler(c *gin.Context) {
	response, err := service.SearchAccessStatistics()
	GenerateResponse(c, response, err)
}

func SearchAccessStatisticsMidHandler(c *gin.Context) {
	response, err := service.SearchAccessStatisticsMid()
	GenerateResponse(c, response, err)
}

func UploadFileHandler(c *gin.Context) {
	fileType, err := strconv.ParseInt(c.Param("type"), 10, 64)
	if err != nil {
		GenerateEmptyDataResponse(c, err)
	}
	header, err := c.FormFile("file")
	if err != nil {
		GenerateEmptyDataResponse(c, err)
	}
	err = service.UploadFile(c, header, consts.FileType(fileType))
	GenerateEmptyDataResponse(c, err)
}
