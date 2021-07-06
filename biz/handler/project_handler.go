package handler

import (
	"github.com/gin-gonic/gin"
	"zhenghui-backend/biz/model/dao"
	"zhenghui-backend/biz/service"
)

func SearchProjectHandler(c *gin.Context) {
	response, err := service.SearchProjectList()
	GenerateResponse(c, response, err)
}

func InsertProjectHandler(c *gin.Context) {
	var project dao.Project
	if err := c.ShouldBind(&project); err != nil {
		GenerateEmptyDataResponse(c, err)
	}
	err := service.InsertProject(project)
	GenerateEmptyDataResponse(c, err)
}

func UpdateProjectHandler(c *gin.Context) {
	var project dao.Project
	if err := c.ShouldBind(&project); err != nil {
		GenerateEmptyDataResponse(c, err)
	}
	err := service.UpateProject(project)
	GenerateEmptyDataResponse(c, err)
}
