package apiserver

import (
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
	app  *fiber.App
	log  *slog.Logger
	addr string
}

func New(service *service.Service, log *slog.Logger, cfg *config.Api) *ApiServer {
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
	controller.Init(app, service, log)

	return &ApiServer{
		app:  app,
		log:  log,
		addr: cfg.Addr,
	}
}
func (a *ApiServer) Run() error {

	chError := make(chan error, 1)
	go func() {
		if err := a.app.Listen(a.addr); err != nil {
			chError <- err
		}
	}()
	go func() {
		chQuit := make(chan os.Signal, 1)
		signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)
		<-chQuit
		chError <- a.app.Shutdown()
	}()

	return <-chError
}
