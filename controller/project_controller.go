package controller

import (
	"github.com/agilistikmal/safatanc-go/model"
	"github.com/agilistikmal/safatanc-go/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ProjectController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type ProjectControllerImpl struct {
	ProjectService service.ProjectService
}

func NewProjectController(projectService service.ProjectService) ProjectControllerImpl {
	return ProjectControllerImpl{ProjectService: projectService}
}

func (controller *ProjectControllerImpl) Create(ctx *fiber.Ctx) error {
	project := new(model.ProjectRequest)
	err := ctx.BodyParser(project)
	if err != nil {
		return ctx.JSON(model.Response{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
		})
	}
	result := controller.ProjectService.Create(project)
	return ctx.JSON(model.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          result,
	})
}

func (controller *ProjectControllerImpl) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	project := new(model.ProjectRequest)
	err := ctx.BodyParser(project)
	if err != nil {
		return ctx.JSON(model.Response{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
		})
	}
	result, err := controller.ProjectService.Update(id, project)
	if err != nil {
		return ctx.JSON(model.Response{
			StatusCode:    http.StatusNotFound,
			StatusMessage: http.StatusText(http.StatusNotFound),
			Data:          err.Error(),
		})
	}
	return ctx.JSON(model.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          result,
	})
}

func (controller *ProjectControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	err := controller.ProjectService.Delete(id)
	if err != nil {
		return ctx.JSON(model.Response{
			StatusCode:    http.StatusNotFound,
			StatusMessage: http.StatusText(http.StatusNotFound),
			Data:          err.Error(),
		})
	}
	return ctx.JSON(model.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	})
}

func (controller *ProjectControllerImpl) FindById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	result, err := controller.ProjectService.FindById(id)
	if err != nil {
		return ctx.JSON(model.Response{
			StatusCode:    http.StatusNotFound,
			StatusMessage: http.StatusText(http.StatusNotFound),
			Data:          err.Error(),
		})
	}
	return ctx.JSON(model.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          result,
	})
}

func (controller *ProjectControllerImpl) FindAll(ctx *fiber.Ctx) error {
	result := controller.ProjectService.FindAll()
	return ctx.JSON(model.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          result,
	})
}
