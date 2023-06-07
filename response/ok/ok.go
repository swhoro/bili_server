package ok

import "github.com/gofiber/fiber/v2"

func ReturnJsonWithOK(c *fiber.Ctx, data fiber.Map) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func ReturnWithSimpleMessage(c *fiber.Ctx, message string) error {
	return ReturnJsonWithOK(c, fiber.Map{"message": message})
}
