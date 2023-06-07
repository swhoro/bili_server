package error

import (
	"github.com/gofiber/fiber/v2"

	"b.carriage.fun/utils/logger"
	"b.carriage.fun/utils/logger/operationCode"
)

type SimpleError struct {
	Message string
}

func (err *SimpleError) Error() string {
	return err.Message
}

// level 1
func ReturnWithErrror(c *fiber.Ctx, errorcode int, message string) error {
	return c.Status(errorcode).JSON(fiber.Map{"message": message})
}

// level 2
func ReturnWithOuterError(c *fiber.Ctx, message string) error {
	return ReturnWithErrror(c, fiber.StatusBadRequest, message)
}

func ReturnWithInternalError(c *fiber.Ctx, err error) error {
	logger.LogInternalError(err)
	return ReturnWithErrror(c, fiber.StatusInternalServerError, "internal error")
}

// level 3
func ReturnWithInvalidInput(c *fiber.Ctx) error {
	return ReturnWithOuterError(c, "invalid input")
}

func ReturnWithNotLogin(c *fiber.Ctx) error {
	return ReturnWithOuterError(c, "not login")
}

func ReturnWithNotAuthorize(c *fiber.Ctx, id uint32, operation operationCode.Operation) error {
	logger.LogNotAuthorizeOperation(id, operation)
	return ReturnWithOuterError(c, "not authorize")
}

func ReturnWithNoUserFound(c *fiber.Ctx) error {
	return ReturnWithOuterError(c, "no user found")
}

func ReturnWithWrongPassword(c *fiber.Ctx) error {
	return ReturnWithOuterError(c, "wrong password")
}
