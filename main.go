package main

import (
	"github.com/agilistikmal/safatanc-go/controller"
	"github.com/agilistikmal/safatanc-go/database"
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
	api := app.Group("/api")

	projectService := service.NewProjectService()
	projectController := controller.NewProjectController(&projectService)

	api.Get("/", projectController.FindAll)
	api.Get("/:id", projectController.FindById)
	api.Post("/", projectController.Create)
	api.Put("/:id", projectController.Update)
	api.Delete("/:id", projectController.Delete)

	err = app.Listen(":8080")
	if err != nil {
		panic(err.Error())
	}
}
