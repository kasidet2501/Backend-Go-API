package logout

import (
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error{


	c.ClearCookie("jwt")

    return c.JSON(fiber.Map{"message": "Logout successful"})
}