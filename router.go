package main

import (
	"github.com/gin-gonic/gin"
	"zhenghui-backend/biz/handler"
)

func customizeRegister(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	root := r.Group("/api/v1")
	{
		root.GET("/get-user-info", handler.SearchUserInfoHandler)
		root.GET("/get-leetcode", handler.SearchLeetcodeHandler)
		root.GET("/get-report", handler.SearchReportHandler)
	}
	rAccess := root.Group("/access")
	{
		rAccess.GET("/insert-data/:page", handler.InsertAccessDataHandler)
		rAccess.GET("/info", handler.SearchAccessHandler)
		rAccess.GET("/statistics", handler.SearchAccessStatisticsHandler)
		rAccess.GET("/statistics-mid", handler.SearchAccessStatisticsMidHandler)
	}
	// Only for test
	r.GET("/test", handler.Test)
}
