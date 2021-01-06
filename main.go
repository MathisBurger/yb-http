package main

import (
	"github.com/MathisBurger/yb-http/config"
	"github.com/MathisBurger/yb-http/installation"
	"github.com/MathisBurger/yb-http/routing"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	installation.Install()
	config.LoadConfigurations()

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/*", routing.Redirect)

	// Start Server
	app.Listen(":80")
}
