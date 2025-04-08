package controller

import (
	"skillsRockTodo/internal/entity"
	repoDto "skillsRockTodo/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	CreateTask(dto *repoDto.DtoCreateTaskReq) error
	GetTasks() ([]*entity.Task, error)
	GetTask(Id int) (*entity.Task, error)
	UpdateTask(dto *repoDto.DtoUpdateTaskReq) error
	DeleteTask(Id int) error
}
type Controller struct {
	service Service
}

func Init(app *fiber.App, service Service) {
	controller := &Controller{
		service: service,
	}
	app.Post("/tasks", controller.CreateTask)
	app.Get("/tasks", controller.GetTasks)
	app.Get("/tasks/:id", controller.GetTasks)
	app.Put("/tasks/:id", controller.UpdateTask)
	app.Delete("/tasks/:id", controller.DeleteTask)
}

func (c *Controller) CreateTask(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)

}
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
func (c *Controller) GetTask(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
func (c *Controller) UpdateTask(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
func (c *Controller) DeleteTask(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
