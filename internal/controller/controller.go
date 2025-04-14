package controller

import (
	"encoding/json"
	"errors"
	"log/slog"
	ctlDto "skillsRockTodo/internal/controller/dto"
	"skillsRockTodo/internal/controller/response"
	"skillsRockTodo/internal/entity"
	repoDto "skillsRockTodo/internal/repository/dto"
	"skillsRockTodo/pkg/servererrors"
	"skillsRockTodo/pkg/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	AddTask(dto *repoDto.AddTask) (*entity.Task, error)
	GetTask(taskId *uuid.UUID) (*entity.Task, error)
	GetTasks(dto *repoDto.GetTasks) ([]*entity.Task, error)
	UpdateTask(dto *repoDto.UpdateTask) (*entity.Task, error)
	RemoveTask(taskId *uuid.UUID) error

	AddUser(dto *repoDto.AddUser) (*entity.User, error)
	GetUser(userId *uuid.UUID) (*entity.User, error)
	GetUsers(dto *repoDto.GetUsers) ([]*entity.User, error)
	UpdateUser(dto *repoDto.UpdateUser) (*entity.User, error)
	RemoveUser(userId *uuid.UUID) error

	AddUserTask(dto *repoDto.AddUserTask) (*entity.UserTask, error)
	GetUserTask(userTaskId *uuid.UUID) (*entity.UserTask, error)
	GetUserTasks(dto *repoDto.GetUserTasks) ([]*entity.UserTask, error)
	RemoveUserTask(userTaskId *uuid.UUID) error
}
type Controller struct {
	service Service
	log     *slog.Logger
}

func Init(app *fiber.App, service Service, log *slog.Logger) {
	controller := &Controller{
		service,
		log,
	}
	app.Post("/tasks", controller.AddTask)
	app.Get("/tasks/:taskId", controller.GetTask)
	app.Get("/tasks", controller.GetTasks)
	app.Patch("/tasks/:taskId", controller.UpdateTask)
	app.Delete("/tasks/:taskId", controller.RemoveTask)
}

func (c *Controller) AddTask(ctx *fiber.Ctx) error {
	var req ctlDto.AddTask
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		c.log.Error("Invalid request body", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.AddTask(&repoDto.AddTask{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		c.log.Error("Failed to insert task", slog.Any("error", err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusCreated(ctx, task)
}
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Queries()["page"])
	if err != nil {
		c.log.Error("Invalid query parameter", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	pageSize, err := strconv.Atoi(ctx.Queries()["pagesize"])
	if err != nil {
		c.log.Error("Invalid query parameter", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	tasks, err := c.service.GetTasks(&repoDto.GetTasks{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		c.log.Error("Failed to get tasks", slog.Any("error", err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, tasks)
}
func (c *Controller) GetTask(ctx *fiber.Ctx) error {
	taskId, err := uuid.Parse(ctx.Params("taskId"))
	if err != nil {
		c.log.Error("Invalid request parameter", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	var req = ctlDto.GetTask{
		TaskId: &taskId,
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.GetTask(req.TaskId)
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		c.log.Error("Failed to get task", slog.Any("error", err))
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		c.log.Error("Failed to get task", slog.Any("error", err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, task)
}
func (c *Controller) UpdateTask(ctx *fiber.Ctx) error {
	var req ctlDto.UpdateTask
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		c.log.Error("Invalid request body", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	taskId, err := uuid.Parse(ctx.Params("taskId"))
	if err != nil {
		c.log.Error("Invalid request parameter", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	req.TaskId = &taskId
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.UpdateTask(&repoDto.UpdateTask{
		TaskId:      req.TaskId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		c.log.Error("Failed to update task", slog.Any("error", err))
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		c.log.Error("Failed to update task", slog.Any("error", err))
		return response.StatusInternalServerError(ctx, err)
	}

	return response.StatusOk(ctx, task)
}
func (c *Controller) RemoveTask(ctx *fiber.Ctx) error {
	taskId, err := uuid.Parse(ctx.Params("taskId"))
	if err != nil {
		c.log.Error("Invalid request parameter", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	var req = ctlDto.GetTask{
		TaskId: &taskId,
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.log.Error("Request data validation error", slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	err = c.service.RemoveTask(req.TaskId)
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		c.log.Error("Failed to delete task", slog.Any("error", err))
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		c.log.Error("Failed to delete task", slog.Any("error", err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusNoContent(ctx)
}
