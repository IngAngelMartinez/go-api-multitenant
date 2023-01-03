package controllers

import (
	"github.com/IngAngelMartinez/go-api-multitenant/src/models"
	"github.com/IngAngelMartinez/go-api-multitenant/src/services"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	documentsService := c.Locals("DocumentsService").(*services.DocumentsService)

	result, err := documentsService.GetAll()

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if result == nil {

		return c.Status(fiber.StatusNotFound).JSON([]models.Document{})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetById(c *fiber.Ctx) error {

	documentsService := c.Locals("DocumentsService").(*services.DocumentsService)

	id := c.Params("id")

	document, err := documentsService.GetById(id)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if document.IsEmpty() {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON(document)
}

func Create(c *fiber.Ctx) error {

	var document models.Document

	if err := c.BodyParser(&document); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	documentsService := c.Locals("DocumentsService").(*services.DocumentsService)

	id, err := documentsService.Create(document)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func Update(c *fiber.Ctx) error {

	documentsService := c.Locals("DocumentsService").(*services.DocumentsService)

	id := c.Params("id")

	var document models.Document

	if err := c.BodyParser(&document); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := documentsService.Update(id, document); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Delete(c *fiber.Ctx) error {

	documentsService := c.Locals("DocumentsService").(*services.DocumentsService)

	id := c.Params("id")

	if err := documentsService.Delete(id); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
