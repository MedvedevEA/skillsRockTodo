package controller

import (
	"log/slog"
	ctrlDto "skillsRockTodo/internal/controller/dto"
	"skillsRockTodo/internal/controller/response"
	repoStoreDto "skillsRockTodo/internal/repository/repostore/dto"
	"skillsRockTodo/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetUsers(ctx *fiber.Ctx) error {
	const op = "controller.GetUsers"
	req := &ctrlDto.GetUsers{
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
	users, err := c.service.GetUsers(&repoStoreDto.GetUsers{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		c.lg.Error("failed to get users", slog.String("op", op), slog.Any("error", err))
		return response.StatusInternalServerError(ctx, err)
	}
	return response.StatusOk(ctx, users)
}
