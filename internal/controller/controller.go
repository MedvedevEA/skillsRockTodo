package controller

import (
	"encoding/json"
	"errors"
	ctlDto "skillsRockTodo/internal/controller/dto"
	"skillsRockTodo/internal/controller/response"
	"skillsRockTodo/internal/entity"
	repoDto "skillsRockTodo/internal/repository/dto"
	"skillsRockTodo/pkg/servererrors"
	"skillsRockTodo/pkg/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Service interface {
	AddTask(dto *repoDto.AddTask) (*entity.Task, error)
	GetTasks() ([]*entity.Task, error)
	GetTask(Id int) (*entity.Task, error)
	UpdateTask(dto *repoDto.UpdateTask) error
	RemoveTask(Id int) error
}
type Controller struct {
	service Service
	log     *zap.SugaredLogger
}

func Init(app *fiber.App, service Service, log *zap.SugaredLogger) {
	controller := &Controller{
		service,
		log,
	}
	app.Post("/tasks", controller.AddTask)
	app.Get("/tasks", controller.GetTasks)
	app.Get("/tasks/:id", controller.GetTask)
	app.Patch("/tasks/:id", controller.UpdateTask)
	app.Delete("/tasks/:id", controller.RemoveTask)
}

func (c *Controller) AddTask(ctx *fiber.Ctx) error {
	var req ctlDto.AddTask
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		c.log.Error("Invalid request body", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.AddTask(&repoDto.AddTask{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		c.log.Error("Failed to insert task", zap.Error(err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusCreated(ctx, task)
}
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := c.service.GetTasks()
	if err != nil {
		c.log.Error("Failed to get tasks", zap.Error(err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, tasks)
}
func (c *Controller) GetTask(ctx *fiber.Ctx) error {
	paramId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		c.log.Error("Invalid request parameter", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	var req = ctlDto.GetTask{
		Id: paramId,
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.GetTask(req.Id)
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		c.log.Error("Failed to get task", zap.Error(err))
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		c.log.Error("Failed to get task", zap.Error(err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, task)
}
func (c *Controller) UpdateTask(ctx *fiber.Ctx) error {
	var req ctlDto.UpdateTask
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		c.log.Error("Invalid request body", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	paramId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		c.log.Error("Invalid request parameter", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	req.Id = paramId
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	err = c.service.UpdateTask(&repoDto.UpdateTask{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		c.log.Error("Failed to update task", zap.Error(err))
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		c.log.Error("Failed to update task", zap.Error(err))
		return response.StatusInternalServerError(ctx, err)
	}

	return response.StatusNoContent(ctx)
}
func (c *Controller) RemoveTask(ctx *fiber.Ctx) error {
	paramId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		c.log.Error("Invalid request parameter", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	var req = ctlDto.GetTask{
		Id: paramId,
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", zap.Error(err))
		return response.StatusBadRequest(ctx, err)
	}
	err = c.service.RemoveTask(req.Id)
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		c.log.Error("Failed to delete task", zap.Error(err))
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		c.log.Error("Failed to delete task", zap.Error(err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusNoContent(ctx)
}
