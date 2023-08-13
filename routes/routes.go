package routes

import (
	"strings"

	"github.com/arzirahman/belajar_golang/handlers"
	"github.com/arzirahman/belajar_golang/utils"
	"github.com/gofiber/fiber/v2"
)

func ValidateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.CreateResponse(c, 401, nil, fiber.Map{"message": "Empty token"})
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return utils.CreateResponse(c, 403, nil, fiber.Map{"message": "Invalid token"})
	}
	claims, err := utils.ValidateAccessToken(tokenParts[1])
	if err != nil {
		return utils.CreateResponse(c, 403, nil, fiber.Map{"message": "Invalid token"})
	}
	c.Locals("claims", claims)
	return c.Next()
}

func SetUserRouter(app *fiber.App) {
	app.Post("/user/login", handlers.UserLogin)
	app.Delete("/user/logout", handlers.UserLogout)
	app.Get("/user/profile", ValidateToken, handlers.UserProfile)

	app.Get("/refresh-token", handlers.RefreshToken)
}
