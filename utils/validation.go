package utils

import (
	"github.com/arzirahman/belajar_golang/models"
	"github.com/gofiber/fiber/v2"
)

func LoginValidation(body *models.Login, c *fiber.Ctx) map[string]string {
	errors := map[string]string{}
	if err := c.BodyParser(&body); err != nil {
		errors["message"] = "Invalid Body Format"
		return errors
	}
	if body.Username == "" {
		errors["username"] = "username is empty"
	}
	if body.Password == "" {
		errors["password"] = "password is empty"
	}
	return errors
}
