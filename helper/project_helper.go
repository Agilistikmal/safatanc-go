package helper

import (
	"github.com/agilistikmal/safatanc-go/model"
	"strings"
)

func ProjectToProjectResponse(project model.Project) model.ProjectResponse {
	techStack := strings.Split(project.TechStack, ",")
	contributors := strings.Split(project.Contributors, ",")
	return model.ProjectResponse{
		Id:           project.Id,
		Title:        project.Title,
		Description:  project.Description,
		TechStack:    techStack,
		Contributors: contributors,
		Image:        project.Image,
		PreviewUrl:   project.PreviewUrl,
		CreatedAt:    project.CreatedAt,
		UpdatedAt:    project.UpdatedAt,
	}
}
