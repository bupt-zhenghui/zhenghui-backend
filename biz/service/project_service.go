package service

import (
	"log"
	"strings"
	"zhenghui-backend/biz/dal/db_dal"
	"zhenghui-backend/biz/model/dao"
	"zhenghui-backend/biz/model/dto"
)

func SearchProjectList() ([]dto.ProjectResponse, error) {
	var response []dto.ProjectResponse
	projects, err := db_dal.SearchProjectList()
	if err != nil {
		log.Println("SearchProjectList error, err = ", err)
		return response, err
	}
	for _, project := range projects {
		res := dto.ProjectResponse{
			Id:         project.Id,
			Type:       project.Type,
			Name:       project.Name,
			Tag:        strings.Split(project.Tag, ","),
			MdURL:      project.MdURL,
			GithubURL:  project.GithubURL,
			ProjectURL: project.ProjectURL,
			CreateTime: project.CreateTime,
			UpdateTime: project.UpdateTime,
		}
		response = append(response, res)
	}
	return response, err
}

func UpateProject(project dao.Project) error {
	return db_dal.UpdateProject(project)
}

func InsertProject(project dao.Project) error {
	return db_dal.InsertProject(project)
}
