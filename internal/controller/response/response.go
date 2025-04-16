package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func StatusBadRequest(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(Response{
		Status: "error",
		Error:  err.Error(),
	})
}

func StatusInternalServerError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status: "error",
		Error:  err.Error(),
	})
}

func StatusUnauthorized(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(Response{
		Status: "error",
		Error:  err.Error(),
	})
}
func StatusForbidden(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusForbidden).JSON(Response{
		Status: "error",
		Error:  err.Error(),
	})
}
func StatusNotFound(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusNotFound).JSON(Response{
		Status: "error",
		Error:  err.Error(),
	})
}
func StatusOk(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Status: "success",
		Data:   data,
	})
}
func StatusCreated(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusCreated).JSON(Response{
		Status: "success",
		Data:   data,
	})
}
func StatusNoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNoContent).JSON(Response{
		Status: "success",
	})
}
