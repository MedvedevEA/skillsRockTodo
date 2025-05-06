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

func (c *Controller) AddTaskUser(ctx *fiber.Ctx) error {
	const op = "controller.AddTaskUser"
	req := new(ctrlDto.AddTaskUser)
	if err := ctx.BodyParser(req); err != nil {
		c.lg.Error("failed to add taskUser", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to add taskUser", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	userTask, err := c.service.AddTaskUser(&repoStoreDto.AddTaskUser{
		UserId: req.UserId,
		TaskId: req.TaskId,
	})
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusCreated(ctx, userTask)
}

func (c *Controller) GetTaskUsers(ctx *fiber.Ctx) error {
	const op = "controller.GetUserTasks"
	req := &ctrlDto.GetTaskUsers{
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
	userTasks, err := c.service.GetTaskUsers(&repoStoreDto.GetTaskUsers{
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
func (c *Controller) RemoveTaskUser(ctx *fiber.Ctx) error {
	const op = "controller.RemoveUserTask"
	req := new(ctrlDto.RemoveTaskUser)
	if err := ctx.ParamsParser(req); err != nil {
		c.lg.Error("failed to remove userTask", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	if err := validator.Validate(ctx.Context(), req); err != nil {
		c.lg.Error("failed to remove userTask", slog.String("op", op), slog.Any("error", err))
		return response.StatusBadRequest(ctx, err)
	}
	err := c.service.RemoveTaskUser(req.TaskUserId)
	if errors.Is(err, servererrors.RecordNotFound) {
		return response.StatusNotFound(ctx, err)
	}
	if err != nil {
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusNoContent(ctx)
}
