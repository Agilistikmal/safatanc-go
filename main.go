package main

import (
	"github.com/agilistikmal/safatanc-go/controller"
	"github.com/agilistikmal/safatanc-go/database"
	"github.com/agilistikmal/safatanc-go/middleware"
	"github.com/agilistikmal/safatanc-go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	database.CreateConnection()

	app := fiber.New()
	projectApi := app.Group("/api/project")
	projectApi.Use(middleware.ProjectMiddleware)

	projectService := service.NewProjectService()
	projectController := controller.NewProjectController(&projectService)

	projectApi.Get("/", projectController.FindAll)
	projectApi.Get("/:id", projectController.FindById)
	projectApi.Post("/", projectController.Create)
	projectApi.Put("/:id", projectController.Update)
	projectApi.Delete("/:id", projectController.Delete)

	err = app.Listen(":8080")
	if err != nil {
		panic(err.Error())
	}
}
