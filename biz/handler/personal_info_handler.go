package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"zhenghui-backend/biz/service"
)

func SearchUserInfoHandler(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		log.Println("Cannot parse user_id to integer type correctly. Error: ", err)
		GenerateEmptyDataResponse(c, err)
	}
	user, err := service.SearchUserInfoById(userId)
	GenerateResponse(c, user, err)
}
