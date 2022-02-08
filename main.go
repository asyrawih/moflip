package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hananloser/moflip/config"
	"github.com/hananloser/moflip/controller"
	"github.com/hananloser/moflip/repository"
	"github.com/hananloser/moflip/services"
)

// Main Function
func main() {

	db := config.Init()
	app := fiber.New()

	userRepo := repository.NewUserRepository(db)
  userService := services.NewUserService(&userRepo)
  userController := controller.NewUserController(&userService)
 
  // Inject App Fiber
  userController.Routes(app)

	// Enable Every Each Request
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Welcome API
	app.Listen(":3000")
}
