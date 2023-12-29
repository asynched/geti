package controllers

import (
	"github.com/asynched/geti/domain/entities"
	"github.com/asynched/geti/domain/repositories"
	"github.com/gofiber/fiber/v2"
)

type VisitController struct {
	linkRepository  repositories.LinkRepository
	visitRepository repositories.VisitRepository
}

func NewVisitController(linkRepository repositories.LinkRepository, visitRepository repositories.VisitRepository) VisitController {
	return VisitController{linkRepository, visitRepository}
}

func (ctrl *VisitController) Create(c *fiber.Ctx) error {
	slug := c.Params("slug")

	link, err := ctrl.linkRepository.FindBySlug(slug)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Link not found",
		})
	}

	_, err = ctrl.visitRepository.Create(entities.Visit{
		Referrer:  c.Get("Referer"),
		UserAgent: c.Get("User-Agent"),
		Ip:        c.IP(),
		LinkId:    link.Id,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating visit",
		})
	}

	return c.Redirect(link.RedirectTo)
}

func (ctrl *VisitController) ListAll(c *fiber.Ctx) error {
	slug := c.Params("slug")

	link, err := ctrl.linkRepository.FindBySlug(slug)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Link not found",
		})
	}

	visits, err := ctrl.visitRepository.FindAll(link.Id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error listing visits",
		})
	}

	return c.JSON(visits)
}
