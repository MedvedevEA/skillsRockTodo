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

	app.Post("/users", controller.AddUser)
	app.Get("/users/:userId", controller.GetUser)
	app.Get("/users", controller.GetUsers)
	app.Patch("/users/:userId", controller.UpdateUser)
	app.Delete("/users/:userId", controller.RemoveUser)

	app.Post("/usertasks", controller.AddUserTask)
	app.Get("/usertasks", controller.GetUserTasks)
	app.Delete("/usertasks/:userTaskId", controller.RemoveUserTask)

}

// Task
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

// User

func (c *Controller) AddUser(ctx *fiber.Ctx) error {
	const op = "controller.AddUser"
	req := new(ctlDto.AddUser)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to add user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to add user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	User, err := c.service.AddUser(&repoDto.AddUser{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusCreated(ctx, User)
}
func (c *Controller) GetUser(ctx *fiber.Ctx) error {
	const op = "controller.GetUser"
	req := new(ctlDto.GetUser)
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to get user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to get user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	user, err := c.service.GetUser(req.UserId)
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, user)
}
func (c *Controller) GetUsers(ctx *fiber.Ctx) error {
	const op = "controller.GetUsers"
	req := &ctlDto.GetUsers{
		Offset: 0,
		Limit:  10,
	}
	if err := ctx.QueryParser(req); err != nil {
		c.lg.Error("failed to get users", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to get users", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	users, err := c.service.GetUsers(&repoDto.GetUsers{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, users)
}
func (c *Controller) UpdateUser(ctx *fiber.Ctx) error {
	const op = "controller.UpdateUser"
	req := new(ctlDto.UpdateUser)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to update user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to update user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to update user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	user, err := c.service.UpdateUser(&repoDto.UpdateUser{
		UserId:   req.UserId,
		Name:     req.Name,
		Password: req.Password,
	})
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, user)
}
func (c *Controller) RemoveUser(ctx *fiber.Ctx) error {
	const op = "controller.RemoveUser"
	req := new(ctlDto.RemoveUser)
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to remove user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to remove user", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	err := c.service.RemoveUser(req.UserId)
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusNoContent(ctx)
}

// UserTask

func (c *Controller) AddUserTask(ctx *fiber.Ctx) error {
	const op = "controller.AddUserTask"
	req := new(ctlDto.AddUserTask)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to add userTask", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to add userTask", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	userTask, err := c.service.AddUserTask(&repoDto.AddUserTask{
		UserId: req.UserId,
		TaskId: req.TaskId,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusCreated(ctx, userTask)
}

func (c *Controller) GetUserTasks(ctx *fiber.Ctx) error {
	const op = "controller.GetUserTasks"
	req := &ctlDto.GetUserTasks{
		Offset: 0,
		Limit:  10,
	}
	if err := ctx.QueryParser(req); err != nil {
		c.lg.Error("failed to get userTasks", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to get userTasks", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	userTasks, err := c.service.GetUserTasks(&repoDto.GetUserTasks{
		Offset: req.Offset,
		Limit:  req.Limit,
		UserId: req.UserId,
		TaskId: req.TaskId,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, userTasks)
}
func (c *Controller) RemoveUserTask(ctx *fiber.Ctx) error {
	const op = "controller.RemoveUserTask"
	req := new(ctlDto.RemoveUserTask)
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to remove userTask", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to remove userTask", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	err := c.service.RemoveUserTask(req.UserTaskId)
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusNoContent(ctx)
}
