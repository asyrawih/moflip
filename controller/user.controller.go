package controller

import (

	"github.com/gofiber/fiber/v2"
	"github.com/hananloser/moflip/entity"
	"github.com/hananloser/moflip/model"
)

type UserController struct {
	app *fiber.App
}

func NewUserController(app *fiber.App) *UserController {
	return &UserController{
		app: app,
	}
}

func (user UserController) Routes() {
	user.app.Get("/users", user.Login)
}
// Login Controller this must depending to Services 
func (user UserController) Login(ctx *fiber.Ctx) error {
	return ctx.JSON(model.WebResponse{
    Code: 200,
    Status: "success",
    Data: entity.User{
      Name: "Hanan Asyrawi",
      Age: 21,
      PhoneNumber: "0862612364162",
      Address: "Tomoni",
      Email: "asyrawih@gmail.com",
    },
  }) 
}
