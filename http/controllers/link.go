package controllers

import (
	"log"

	"github.com/asynched/geti/domain/repositories"
	"github.com/asynched/geti/dto"
	"github.com/gofiber/fiber/v2"
)

type LinkController struct {
	linkRepository repositories.LinkRepository
}

func NewLinkController(linkRepository repositories.LinkRepository) LinkController {
	return LinkController{linkRepository}
}

func (ctrl *LinkController) Create(c *fiber.Ctx) error {
	data := dto.CreateLink{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := data.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	link, err := ctrl.linkRepository.Create(data.ToEntity())

	if err != nil {
		log.Println("Error creating link:", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating link",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(link)
}
