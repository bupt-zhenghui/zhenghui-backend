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
		root.GET("/get-navigation")
	}
	// Only for test
	//r.GET("/test", handler.Test)
}
