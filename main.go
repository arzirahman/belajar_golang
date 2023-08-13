package main

import (
	"fmt"
	"os"

	"github.com/arzirahman/belajar_golang/routes"
	"github.com/arzirahman/belajar_golang/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	utils.SetDbConnection()

	routes.SetUserRouter(app)

	if appPort := os.Getenv("APP_PORT"); appPort == "" {
		app.Listen(":5000")
	} else {
		app.Listen(":" + appPort)
	}
}
