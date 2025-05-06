package apiserver

import (
	"crypto/rsa"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"skillsRockTodo/internal/config"
	"skillsRockTodo/internal/controller"
	"skillsRockTodo/pkg/secure"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ApiServer struct {
	app *fiber.App
	lg  *slog.Logger
	cfg *config.Api
}
type TokenClaims struct {
	Jti        *uuid.UUID `json:"jti"`
	Sub        *uuid.UUID `json:"sub"`
	DeviceCode string     `json:"device"`
	TokenType  string     `json:"type"`
	jwt.RegisteredClaims
}

func MustNew(ctrl *controller.Controller, lg *slog.Logger, cfg *config.Api) *ApiServer {
	const op = "apiserver.MustNew"
	publicKey, err := secure.LoadPublicKey(cfg.PublicKeyPath)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}

	app := fiber.New(fiber.Config{
		AppName:      cfg.Name,
		WriteTimeout: cfg.WriteTimeout,
	})

	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-SomeID",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))
	app.Use(recover.New(recover.ConfigDefault))

	app.Use(GetLoggerMiddlewareFunc(lg))
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	apiGroup := app.Group("/api")
	v1Group := apiGroup.Group("/v1")
	//public
	publicGroup := v1Group.Group("")
	publicGroup.Post("/login", func(ctx *fiber.Ctx) error { return ctx.SendStatus(fiber.StatusNotImplemented) })
	publicGroup.Post("/registration", func(ctx *fiber.Ctx) error { return ctx.SendStatus(fiber.StatusNotImplemented) })
	//refresh
	refreshGroup := v1Group.Group("")
	refreshGroup.Use(keyauth.New(keyauth.Config{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator:  GetTokenValidateFunc(publicKey, "refresh"),
	}))
	refreshGroup.Post("/refresh", func(ctx *fiber.Ctx) error { return ctx.SendStatus(fiber.StatusNotImplemented) })
	//access
	accessGroup := v1Group.Group("")
	accessGroup.Use(keyauth.New(keyauth.Config{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator:  GetTokenValidateFunc(publicKey, "access"),
	}))
	accessGroup.Post("/logout", func(ctx *fiber.Ctx) error { return ctx.SendStatus(fiber.StatusNotImplemented) })
	accessGroup.Post("/unregistration", func(ctx *fiber.Ctx) error { return ctx.SendStatus(fiber.StatusNotImplemented) })

	accessGroup.Post("/messages", ctrl.AddMessage)
	accessGroup.Get("/messages/:messagesId", ctrl.GetMessage)
	accessGroup.Get("/messages", ctrl.GetMessages)
	accessGroup.Patch("/messages/:messageId", ctrl.UpdateMessage)
	accessGroup.Delete("/messages/:messageId", ctrl.RemoveMessage)

	accessGroup.Post("/statuses", ctrl.AddStatus)
	accessGroup.Get("/statuses/:markId", ctrl.GetStatus)
	accessGroup.Get("/statuses", ctrl.GetStatuses)
	accessGroup.Patch("/statuses/:statusId", ctrl.UpdateStatus)
	accessGroup.Delete("/statuses/:statusId", ctrl.RemoveStatus)

	accessGroup.Post("/tasks", ctrl.AddTask)
	accessGroup.Get("/tasks/:taskId", ctrl.GetTask)
	accessGroup.Get("/tasks", ctrl.GetTasks)
	accessGroup.Patch("/tasks/:taskId", ctrl.UpdateTask)
	accessGroup.Delete("/tasks/:taskId", ctrl.RemoveTask)

	accessGroup.Post("/taskusers", ctrl.AddTaskUser)
	accessGroup.Get("/taskusers", ctrl.GetTaskUsers)
	accessGroup.Delete("/taskusers/:taskUserId", ctrl.RemoveTaskUser)

	accessGroup.Get("/users", ctrl.GetUsers)

	return &ApiServer{
		app,
		lg,
		cfg,
	}
}
func GetLoggerMiddlewareFunc(lg *slog.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		lg.Info(
			"api server request",
			slog.Any("method", c.Method()),
			slog.Any("path", c.Path()),
			slog.Any("status", c.Response().StatusCode()),
			slog.Any("time", time.Since(start)),
		)

		return err

	}
}

func GetTokenValidateFunc(publicKey *rsa.PublicKey, tokenType string) func(c *fiber.Ctx, tokenString string) (bool, error) {
	return func(c *fiber.Ctx, tokenString string) (bool, error) {
		tokenJwt, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
		if err != nil {
			return false, err
		}
		tokenClaims, ok := tokenJwt.Claims.(*TokenClaims)
		if !ok {
			return false, errors.New("get token claims error")
		}
		if tokenClaims.TokenType != tokenType {
			return false, errors.New("invalid token type")
		}
		c.Locals("claims", tokenClaims)
		return true, nil
	}
}
func (a *ApiServer) Run() {

	const op = "apiserver.Run"
	chError := make(chan error, 1)
	go func() {
		a.lg.Info(fmt.Sprintf("API Server '%s' is started", a.cfg.Name), slog.String("op", op), slog.String("bind address", a.cfg.Addr))
		if err := a.app.Listen(a.cfg.Addr); err != nil {
			chError <- err
		}
	}()
	go func() {
		chQuit := make(chan os.Signal, 1)
		signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)
		<-chQuit
		chError <- a.app.Shutdown()
	}()
	if err := <-chError; err != nil {
		a.lg.Error(fmt.Sprintf("API Server '%s' error", a.cfg.Name), slog.String("op", op), slog.Any("error", err))
		return
	}
	a.lg.Info(fmt.Sprintf("API Server '%s' is stopped", a.cfg.Name), slog.String("op", op))

}
