package middleware

import (
	"Tahagram/httpstatus"
	"Tahagram/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authorized, _ := auth.AuthenticateUser(c)

	if authorized {
		c.Next()
	} else {
		httpstatus.Unauthorized(c)
	}

	return nil
}
