package apiserver

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"skillsRockTodo/internal/apiserver/middleware"
	"skillsRockTodo/internal/config"
	"skillsRockTodo/internal/controller"
	"skillsRockTodo/internal/service"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ApiServer struct {
	app *fiber.App
	lg  *slog.Logger
	cfg *config.Api
}

func New(service *service.Service, lg *slog.Logger, cfg *config.Api) *ApiServer {
	app := fiber.New(fiber.Config{
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
	app.Use(middleware.Authorization())
	controller.Init(app, service, lg)

	return &ApiServer{
		app,
		lg,
		cfg,
	}
}
func (a *ApiServer) MustRun() {

	const op = "apiserver.Run"
	chError := make(chan error, 1)
	go func() {
		a.lg.Info(fmt.Sprintf("API Server '%s' is started in addr:[%s]", a.cfg.Name, a.cfg.Addr), slog.String("op", op))
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
