package utils

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message any    `json:"message"`
}

type FailedResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors any    `json:"errors"`
}

func CreateResponse(c *fiber.Ctx, code int, data any, errors any) error {
	var response any
	if code == fiber.StatusOK {
		response = SuccessResponse{
			Code:    code,
			Status:  getStatus(code),
			Message: data,
		}
		return c.Status(code).JSON(response)
	}
	response = FailedResponse{
		Code:   code,
		Status: getStatus(code),
		Errors: errors,
	}
	return c.Status(code).JSON(response)
}

func getStatus(code int) string {
	switch code {
	case fiber.StatusOK:
		return "OK"
	case fiber.StatusBadRequest:
		return "Bad Request"
	case fiber.StatusUnauthorized:
		return "Unauthrized"
	case fiber.StatusForbidden:
		return "Forbidden"
	case fiber.StatusInternalServerError:
		return "Internal Server Error"
	default:
		return ""
	}
}
