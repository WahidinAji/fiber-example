package helpers

import (
	"github.com/gofiber/fiber/v2"
)
func ResponseJson(ctx *fiber.Ctx, code int, message string, status bool, data interface{}) error {
	response := ctx.Status(code).JSON(&fiber.Map{
		"code": ctx.Response().StatusCode(),
		"message": message,
		"status": status,
		"data": data,
	})
	ctx.Response().Header.SetContentType("application/json")
	return response
}

type WebResponse struct{
	Code int `json:"code"`
	Status string    `json:"status"`
	Data interface{} `json:"data"`
}