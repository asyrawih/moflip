package helper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandlerJwt(ctx *fiber.Ctx, err error) error {
	return ctx.JSON(fiber.Map{
		"code":  "Malformat JWT",
		"error": err.Error(),
	})
}

func SuccessHandler(ctx *fiber.Ctx) error {
	fmt.Println("Hitting in here when code is success")
	return ctx.Next()
}
