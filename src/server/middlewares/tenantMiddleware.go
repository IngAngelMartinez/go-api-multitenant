package middlewares

import (
	"fmt"

	"github.com/IngAngelMartinez/go-api-multitenant/src/dependencies"
	"github.com/gofiber/fiber/v2"
)

func TenantContext() fiber.Handler {

	return func(c *fiber.Ctx) error {

		tenantId := c.Get("TenantId")

		if tenantId == "" {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "tenantId not set"})
		}

		tenantDependecies, err := dependencies.GetTenantDependencies(tenantId)

		if err != nil {

			message := fmt.Sprintf("tenantId '%s' not found", tenantId)

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": message})
		}

		c.Locals("DocumentsService", tenantDependecies.DocumentsService)

		return c.Next()
	}
}
