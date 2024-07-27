package handler

import (
	"github.com/abyan-dev/productivity/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func StartPomodoro(c *fiber.Ctx) error {
	return response.Ok(c, "")
}

func StopPomodoro(c *fiber.Ctx) error {
	return response.Ok(c, "")
}
