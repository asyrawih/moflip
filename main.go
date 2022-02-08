package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hananloser/moflip/model"
)

// Main Function
func main() {

	app := fiber.New()
	// Enable Every Each Request
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

  // Welcome API
	app.Get("/", func(c *fiber.Ctx) error {
    return c.JSON(model.WebResponse{
      Code: 200,
      Status: "success",
      Data: "Moflip API",
    })
	})

	app.Listen(":3000")
}
