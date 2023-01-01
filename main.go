package main

import (
	"github.com/IngAngelMartinez/go-api-multitenant/src/configurations"
	"github.com/IngAngelMartinez/go-api-multitenant/src/server"
	"github.com/IngAngelMartinez/go-api-multitenant/src/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fiberConfigurations := configurations.GetServerConfigurations()

	app := fiber.New(fiberConfigurations)

	app.Use(middlewares.TenantContext())

	server.MapRoutes(app)

	app.Listen(":3000")
}
