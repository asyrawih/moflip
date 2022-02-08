package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hananloser/moflip/helper"
	"github.com/hananloser/moflip/middleware"
	"github.com/hananloser/moflip/model"
	"github.com/hananloser/moflip/services"
)

type UserController interface {
	Index(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Routes(app *fiber.App)
}

func NewUserController(userService *services.UserService) UserController {
	return &UserControllerImpl{
		userService: *userService,
	}
}

// UserControllerImpl struct  
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

// Login method    login
func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	user := model.UserLoginModel{}
	err := ctx.BodyParser(&user)
	helper.PanicIfNeeded(err)
	result := controller.userService.FindUsernameAndPassword(user.Email, user.Password)

	claims := jwt.MapClaims{
		"name":  result.Name,
		"email": result.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create Intance JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Encode The Token
	t, err := token.SignedString([]byte("secret"))

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "Success",
		Data: fiber.Map{
			"email": result.Email,
			"name":  result.Name,
			"token": t,
		},
	})
}

// Define Routes For Each Method
func (controller *UserControllerImpl) Routes(app *fiber.App) {
	app.Post("/users/login", controller.Login)
	// Protected the Routes
	app.Get("/users", middleware.Protected(), controller.Index)
}
