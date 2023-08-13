package handlers

import (
	"github.com/arzirahman/belajar_golang/models"
	"github.com/arzirahman/belajar_golang/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *fiber.Ctx) error {
	var body models.Login
	if errors := utils.LoginValidation(&body, c); len(errors) > 0 {
		return utils.CreateResponse(c, 400, nil, errors)
	}
	var users []models.User
	result := utils.DB.Raw(`SELECT * FROM "User" WHERE "username" = ?`, body.Username).Scan(&users)
	if result.Error != nil {
		return utils.CreateResponse(c, 500, nil, fiber.Map{"message": "Login failed"})
	}
	if len(users) == 0 {
		return utils.CreateResponse(c, 400, nil, fiber.Map{"message": "Invalid username or password"})
	}
	err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(body.Password))
	if err != nil {
		return utils.CreateResponse(c, 400, nil, fiber.Map{"message": "Invalid username or password"})
	}
	token := utils.GenerateAccessToken(jwt.MapClaims{"username": users[0].Username})
	return utils.CreateResponse(c, 200, fiber.Map{"message": token}, nil)
}

func UserProfile(c *fiber.Ctx) error {
	return utils.CreateResponse(c, 200, fiber.Map{"message": "lol"}, nil)
}
