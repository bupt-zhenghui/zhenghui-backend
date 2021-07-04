package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"zhenghui-backend/biz/consts"
	"zhenghui-backend/biz/model/dao"
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

type Calendar struct {
	Date string `gorm:"column:datelist" json:"datelist"`
}

func Synchronize(c *gin.Context) {
	stamp := time.Now().AddDate(0, 0, -2)
	for i := 0; i < 2000; i++ {
		cur := stamp.Format(consts.FormatDate)
		date := Calendar{
			Date: cur,
		}
		if err := dao.DB.Create(date).Error; err != nil {
			GenerateEmptyDataResponse(c, err)
		}
		stamp = stamp.AddDate(0, 0, 1)
	}
	GenerateEmptyDataResponse(c, nil)
}
