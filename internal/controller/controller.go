package controller

import (
	"errors"
	"log/slog"
	ctlDto "skillsRockTodo/internal/controller/dto"
	"skillsRockTodo/internal/controller/response"
	"skillsRockTodo/internal/entity"
	repoDto "skillsRockTodo/internal/repository/dto"
	"skillsRockTodo/pkg/servererrors"
	"skillsRockTodo/pkg/validator"

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
	lg      *slog.Logger
}

func Init(app *fiber.App, service Service, lg *slog.Logger) {
	controller := &Controller{
		service,
		lg,
	}
	app.Post("/tasks", controller.AddTask)
	app.Get("/tasks/:taskId", controller.GetTask)
	app.Get("/tasks", controller.GetTasks)
	app.Patch("/tasks/:taskId", controller.UpdateTask)
	app.Delete("/tasks/:taskId", controller.RemoveTask)
}

func (c *Controller) AddTask(ctx *fiber.Ctx) error {
	const op = "controller.AddTask"
	req := new(ctlDto.AddTask)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to add task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to add task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.AddTask(&repoDto.AddTask{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusCreated(ctx, task)
}
func (c *Controller) GetTask(ctx *fiber.Ctx) error {
	const op = "controller.GetTask"
	req := new(ctlDto.GetTask)
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to get task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to get task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.GetTask(req.TaskId)
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, task)
}
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	const op = "controller.GetTasks"
	req := &ctlDto.GetTasks{
		Offset: 0,
		Limit:  10,
	}
	if err := ctx.QueryParser(req); err != nil {
		c.lg.Error("failed to get tasks", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to get tasks", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	tasks, err := c.service.GetTasks(&repoDto.GetTasks{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, tasks)
}
func (c *Controller) UpdateTask(ctx *fiber.Ctx) error {
	const op = "controller.UpdateTask"
	req := new(ctlDto.UpdateTask)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to update task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to update task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to update task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.UpdateTask(&repoDto.UpdateTask{
		TaskId:      req.TaskId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, task)
}
func (c *Controller) RemoveTask(ctx *fiber.Ctx) error {
	const op = "controller.RemoveTask"
	req := new(ctlDto.RemoveTask)
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to remove task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to remove task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	err := c.service.RemoveTask(req.TaskId)
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusNoContent(ctx)
}
