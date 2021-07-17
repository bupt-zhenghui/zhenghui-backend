package db_dal

import (
	"zhenghui-backend/biz/model/dao"
)

func SearchProjectList() ([]dao.Project, error) {
	var response []dao.Project
	err := dao.DB.Find(&response).Error
	return response, err
}

func InsertProject(project dao.Project) error {
	return dao.DB.Create(&project).Error
}

func UpdateProject(project dao.Project) error {
	return dao.DB.Table("project").Where("id = ?", project.Id).Update("name", project.Name).Update("tag", project.Tag).
		Update("md_url", project.MdURL).Update("github_url", project.GithubURL).Update("project_url", project.ProjectURL).Error
}

func DeleteProject(id int64) error {
	return dao.DB.Delete(&dao.Project{}, id).Error
}
