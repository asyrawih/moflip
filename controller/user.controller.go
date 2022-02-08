package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hananloser/moflip/model"
	"github.com/hananloser/moflip/services"
)

type UserController interface {
	Index(ctx *fiber.Ctx) error

	Routes(app *fiber.App)
}

func NewUserController(userService *services.UserService) UserController {
	return &UserControllerImpl{
		userService: *userService,
	}
}

type UserControllerImpl struct {
	userService services.UserService
}

func (controller *UserControllerImpl) Index(ctx *fiber.Ctx) error {

	user := controller.userService.GetAllUser()

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "Success Getted Data",
		Data:   user,
	})

}

func (controller *UserControllerImpl) Routes(app *fiber.App) {
	app.Get("/users", controller.Index)
}
