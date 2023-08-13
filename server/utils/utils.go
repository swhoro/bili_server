package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"b.carriage.fun/server/global"
	error2 "b.carriage.fun/server/response/error"
)

func GetInformationFromJWTToken(c *fiber.Ctx, param string) (any, error) {
	str := c.Locals("jwt_token")
	if str == nil {
		return "", &error2.SimpleError{Message: "no jwt token found"}
	}
	token := str.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	result := claims[param]
	return result, nil
}

func CreateJWTToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(global.Secret))
	if err != nil {
		return "", &error2.SimpleError{Message: "internal server error"}
	}
	return t, nil
}
