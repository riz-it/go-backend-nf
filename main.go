package main

import (
	"fmt"
	"log"
	"strconv"

	"riz.it/nurul-faizah/internal/config"
	"riz.it/nurul-faizah/internal/injector"
)

func main() {
	// cnf := config.Get()
	// log := config.NewLogger(cnf)
	// db := config.NewDatabase(cnf, log)
	// validate := config.NewValidator(cnf)
	// app := config.NewFiber(cnf)
	app := injector.InitializedApp()
	cnf := config.Get()

	// config.Bootstrap(&config.BootstrapConfig{
	// 	DB:       db,
	// 	App:      app,
	// 	Log:      log,
	// 	Validate: validate,
	// })
	port, _ := strconv.Atoi(cnf.Server.Port)
	err := app.Fiber.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
