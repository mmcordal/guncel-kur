package main

import (
	"guncel-kur/internal/infrastructure/app"
	"guncel-kur/internal/router"
)

func main() {
	r := router.NewRouter()
	a := app.New(r)
	a.Start()
}
