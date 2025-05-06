package controller

import (
	"errors"
	"log/slog"
	ctrlDto "skillsRockTodo/internal/controller/dto"
	"skillsRockTodo/internal/controller/response"
	repoStoreDto "skillsRockTodo/internal/repository/repostore/dto"
	"skillsRockTodo/pkg/servererrors"
	"skillsRockTodo/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) AddTask(ctx *fiber.Ctx) error {

	const op = "controller.AddTask"
	req := new(ctrlDto.AddTask)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to add task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to add task", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	task, err := c.service.AddTask(&repoStoreDto.AddTask{
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
	req := new(ctrlDto.GetTask)
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
	req := &ctrlDto.GetTasks{
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
	tasks, err := c.service.GetTasks(&repoStoreDto.GetTasks{
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
	req := new(ctrlDto.UpdateTask)
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
	task, err := c.service.UpdateTask(&repoStoreDto.UpdateTask{
		TaskId:      req.TaskId,
		StatusId:    req.StatusId,
		Title:       req.Title,
		Description: req.Description,
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
	req := new(ctrlDto.RemoveTask)
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
