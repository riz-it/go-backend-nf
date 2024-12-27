package config

import (
	"github.com/gofiber/fiber/v2"
	"riz.it/nurul-faizah/internal/delivery/route"
)

type App struct {
	Fiber  *fiber.App
	Config *Bootstrap
}

func NewApp(
	fiber *route.RouterConfig,
	config *Bootstrap,
) *App {
	return &App{
		Fiber:  fiber.App,
		Config: config,
	}
}
