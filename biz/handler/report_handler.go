package handler

import (
	"github.com/gin-gonic/gin"
	"zhenghui-backend/biz/service"
)

func SearchReportHandler(c *gin.Context) {
	reportList, err := service.SearchMonthlyReport()
	GenerateResponse(c, reportList, err)
}
