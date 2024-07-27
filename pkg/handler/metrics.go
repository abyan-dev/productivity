package handler

import (
	"github.com/abyan-dev/productivity/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func GenerateStudyMetrics(c *fiber.Ctx) error {
	return response.Ok(c, "")
}
