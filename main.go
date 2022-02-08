package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hananloser/moflip/controller"
	"github.com/hananloser/moflip/model"
)

// Main Function
func main() {

	app := fiber.New()

	// Enable Every Each Request
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// userController
	userController := controller.NewUserController(app)
	userController.Routes()

	// Welcome API
	app.Get("/", Welcome)

	app.Listen(":3000")
}

func Welcome(ctx *fiber.Ctx) error {
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   "Moflip API",
	})
}
