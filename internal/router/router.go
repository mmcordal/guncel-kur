package router

import (
	"guncel-kur/internal/handler"
	"guncel-kur/internal/infrastructure/app"
	"guncel-kur/internal/service"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (Router) RegisterRouter(a *app.App) {
	app := a.FiberApp
	rds := a.Redis

	kurService := service.NewKurService(rds)

	kurHandler := handler.NewKurHandler(kurService)

	v1 := app.Group("/api/v1")

	v1.Get("/guncel-kur", kurHandler.GetKur)
}
