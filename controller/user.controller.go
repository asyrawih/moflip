package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
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

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	user := model.UserLoginModel{}
	err := ctx.BodyParser(&user)

	if err != nil {
		fmt.Println(err)
	}
	result := controller.userService.FindUsernameAndPassword(user.Email, user.Password)

	fmt.Println(result)

	claims := jwt.MapClaims{
		"name":  result.Name,
		"email": result.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Generate en// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	fmt.Println("Token in Here" + t)

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

// Define Routes For Each Controller
func (controller *UserControllerImpl) Routes(app *fiber.App) {
	app.Post("/users/login", controller.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/users", controller.Index)
}
