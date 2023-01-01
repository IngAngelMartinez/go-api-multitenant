package server

import (
	"github.com/IngAngelMartinez/go-api-multitenant/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {

	apiRouter := app.Group("/api")

	documentsRoute(apiRouter)
}

func documentsRoute(apiRouter fiber.Router) {

	apiRouter.Get("/documents", controllers.GetAll)

	apiRouter.Get("/documents/:id", controllers.GetById)

	apiRouter.Post("/documents/", controllers.Create)
}
