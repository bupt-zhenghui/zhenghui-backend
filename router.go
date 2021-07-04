package main

import (
	"github.com/gin-gonic/gin"
	"zhenghui-backend/biz/handler"
)

func customizeRegister(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	root := r.Group("/api/v1")
	{
		root.GET("/insert-access-data/:page", handler.InsertAccessDataHandler)
		root.GET("/get-access-info", handler.SearchAccessHandler)
		root.GET("/get-access-statistics", handler.SearchAccessStatisticsHandler)
		root.GET("/get-user-info", handler.SearchUserInfoHandler)
		root.GET("/get-leetcode", handler.SearchLeetcodeHandler)
		root.GET("/get-report", handler.SearchReportHandler)
	}
	// Only for test
	r.GET("/test", handler.Test)
}
