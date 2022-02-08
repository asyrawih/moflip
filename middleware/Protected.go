package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/hananloser/moflip/helper"
)

// Protected function    
// Handle The Authenticated Routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte("secret"),
		SuccessHandler: helper.SuccessHandler,
		ErrorHandler:   helper.ErrorHandlerJwt,
	})
}
