package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type HealthController struct {
	startup time.Time
}

func NewHealthController() HealthController {
	return HealthController{time.Now()}
}

func (ctrl *HealthController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "up",
		"uptime": time.Since(ctrl.startup).String(),
	})
}
