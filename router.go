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
		root.POST("/upload/:type", handler.UploadFileHandler)
	}
	rAccess := root.Group("/access")
	{
		rAccess.GET("/insert-data/:page", handler.InsertAccessDataHandler)
		rAccess.GET("/info", handler.SearchAccessHandler)
		rAccess.GET("/statistics", handler.SearchAccessStatisticsHandler)
		rAccess.GET("/statistics-mid", handler.SearchAccessStatisticsMidHandler)
	}
	rProject := root.Group("/project")
	{
		rProject.GET("/search", handler.SearchProjectHandler)
		rProject.POST("/create", handler.InsertProjectHandler)
		rProject.POST("/update", handler.UpdateProjectHandler)
		rProject.GET("/delete/:id", handler.DeleteProjectHandler)
	}
	// Only for test
	r.GET("/test", handler.Test)
}
