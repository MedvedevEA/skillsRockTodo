package apiserver

import (
	"os"
	"os/signal"
	"skillsRockTodo/internal/apiserver/middleware"
	"skillsRockTodo/internal/config"
	"skillsRockTodo/internal/controller"
	"skillsRockTodo/internal/service"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type ApiServer struct {
	server        *fiber.App
	log           *zap.SugaredLogger
	ListenAddress string
}

func New(service *service.Service, log *zap.SugaredLogger, cfg config.Rest) *ApiServer {
	app := fiber.New(fiber.Config{
		WriteTimeout: cfg.WriteTimeout,
	})
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-SomeID",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))
	app.Use(recover.New(recover.ConfigDefault))
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(middleware.Authorization(cfg.Token))
	controller.Init(app, service, log)

	return &ApiServer{
		server:        app,
		log:           log,
		ListenAddress: cfg.ListenAddress,
	}
}
func (a *ApiServer) Run() error {

	chError := make(chan error, 1)
	go func() {
		if err := a.server.Listen(a.ListenAddress); err != nil {
			chError <- err
		}
	}()
	go func() {
		chQuit := make(chan os.Signal, 1)
		signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)
		<-chQuit
		chError <- a.server.Shutdown()
	}()

	return <-chError
}
