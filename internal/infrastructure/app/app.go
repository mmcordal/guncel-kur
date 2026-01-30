package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"guncel-kur/internal/infrastructure/cache"
	"guncel-kur/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	FiberApp *fiber.App
	Redis    *cache.RedisClient
	Cfg      *config.Config
}

type IRouter interface {
	RegisterRouter(app *App)
}

func New(router IRouter) *App {

	cfg, err := config.Setup()
	if err != nil {
		panic(err)
	}
	fiberApp := fiber.New()

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // http://localhost:5173	http://---IP---:5173
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	redisClient := cache.NewRedisClient(
		cfg.Redis.Host + ":" + cfg.Redis.Port,
	)

	app := &App{
		FiberApp: fiberApp,
		Redis:    redisClient,
		Cfg:      cfg,
	}

	router.RegisterRouter(app)

	return app
}

func (a *App) Start() {
	go func() {
		err := a.FiberApp.Listen(fmt.Sprintf(":%v", a.Cfg.Server.Port))
		if err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
}
