package handlers

import (
	"github.com/arzirahman/belajar_golang/utils"
	"github.com/gofiber/fiber/v2"
)

func RefreshToken(c *fiber.Ctx) error {
	token := c.Cookies("session")
	if token == "" {
		return utils.CreateResponse(c, 401, nil, fiber.Map{"message": "Empty session"})
	}
	claims, err := utils.ValidateRefreshToken(token)
	if err != nil {
		return utils.CreateResponse(c, 401, nil, fiber.Map{"message": "Invalid session"})
	}
	accessToken := utils.GenerateAccessToken(claims)
	return utils.CreateResponse(c, 200, fiber.Map{"message": accessToken}, nil)
}
