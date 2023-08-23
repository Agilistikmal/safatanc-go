package service

import (
	"fmt"
	"github.com/agilistikmal/safatanc-go/database"
	"github.com/agilistikmal/safatanc-go/helper"
	"github.com/agilistikmal/safatanc-go/model"
	"strings"
)

type ProjectService interface {
	Create(project *model.ProjectRequest) model.ProjectResponse
	Update(projectId int, project *model.ProjectRequest) (model.ProjectResponse, error)
	Delete(projectId int) error
	FindById(projectId int) (model.ProjectResponse, error)
	FindAll(page int, limit int) ([]model.ProjectResponse, int)
}

type ProjectServiceImpl struct {
}

func NewProjectService() ProjectServiceImpl {
	return ProjectServiceImpl{}
}

func (service *ProjectServiceImpl) Create(request *model.ProjectRequest) model.ProjectResponse {
	techStack := strings.Join(request.TechStack, ",")
	contributors := strings.Join(request.Contributors, ",")
	project := model.Project{
		Title:        request.Title,
		Description:  request.Description,
		TechStack:    techStack,
		Contributors: contributors,
		Image:        request.Image,
		PreviewUrl:   request.PreviewUrl,
	}
	database.DB.Create(&project)
	return helper.ProjectToProjectResponse(project)
}

func (service *ProjectServiceImpl) Update(projectId int, request *model.ProjectRequest) (model.ProjectResponse, error) {
	_, err := service.FindById(projectId)
	if err != nil {
		return model.ProjectResponse{}, err
	}
	var project model.Project
	database.DB.First(&project, "id = ?", projectId)

	techStack := strings.Join(request.TechStack, ",")
	contributors := strings.Join(request.Contributors, ",")

	project.Title = request.Title
	project.Image = request.Image
	project.TechStack = techStack
	project.Contributors = contributors
	project.Description = request.Description
	project.PreviewUrl = request.PreviewUrl
	database.DB.Save(&project)

	return helper.ProjectToProjectResponse(project), nil

}

func (service *ProjectServiceImpl) Delete(projectId int) error {
	_, err := service.FindById(projectId)
	if err != nil {
		return err
	}
	var project model.Project
	database.DB.Where("id = ?", projectId).Delete(&project)
	return nil
}

func (service *ProjectServiceImpl) FindById(projectId int) (model.ProjectResponse, error) {
	var project model.Project
	result := database.DB.First(&project, "id = ?", projectId)
	if result.Error != nil {
		return model.ProjectResponse{}, fmt.Errorf("project with id %d not found", projectId)
	}
	return helper.ProjectToProjectResponse(project), nil
}

func (service *ProjectServiceImpl) FindAll(page int, limit int) ([]model.ProjectResponse, int) {
	var projects []model.Project
	var responses []model.ProjectResponse
	var count int64

	offset := (page - 1) * limit
	database.DB.Offset(offset).Limit(limit).Find(&projects)
	database.DB.Table("projects").Count(&count)
	totalPage := (int(count) / limit) + 1
	for _, project := range projects {
		responses = append(responses, helper.ProjectToProjectResponse(project))
	}
	return responses, totalPage
}
