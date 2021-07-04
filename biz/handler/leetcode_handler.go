package handler

import (
	"github.com/gin-gonic/gin"
	"zhenghui-backend/biz/service"
)

func SearchLeetcodeHandler(c *gin.Context) {
	codeList, err := service.SearchLeetcode()
	GenerateResponse(c, codeList, err)
}
