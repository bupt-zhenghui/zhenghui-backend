/*
Author: Zhenghui Wang
Affiliation: BUPT, Bytedance Inc
Revised Time: 2021-06-06
Version: 1.0.0
All Rights Reserved
*/

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"zhenghui-backend/biz/middleware"
	"zhenghui-backend/biz/model/dao"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	dao.ConnectDB()
	customizeRegister(r)
	if err := r.Run(); err != nil {
		log.Fatal("Initialize server failed, error: ", err)
	}
}
