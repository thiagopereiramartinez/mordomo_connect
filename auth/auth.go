package auth

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func Auth(c *fiber.Ctx) {
	redirect_uri, state := c.Query("redirect_uri"), c.Query("state")

	c.Redirect(fmt.Sprintf("%s?code=%s&state=%s", redirect_uri, uuid.New().String(), state))
}

func Token(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"token_type":    "Bearer",
		"access_token":  uuid.New().String(),
		"refresh_token": uuid.New().String(),
		"expires_in":    86400,
	})
}
