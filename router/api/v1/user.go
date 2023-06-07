package v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	g "b.carriage.fun/global"
	"b.carriage.fun/model"
	responseError "b.carriage.fun/response/error"
	responseOK "b.carriage.fun/response/ok"
	"b.carriage.fun/utils"
	"b.carriage.fun/utils/logger/operationCode"
)

// AddUser 新增用户 POST
// 传入格式:
// username: 用户名
// password: 用户名
func AddUser(c *fiber.Ctx) error {
	type AddUserInput struct {
		Username string
		Password string
	}

	in := new(AddUserInput)
	err := c.BodyParser(in)
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	username := in.Username
	password := in.Password
	if username == "" || password == "" {
		return responseError.ReturnWithInvalidInput(c)
	}

	user := new(model.User)
	err = g.DB.First(user, "username = ?", username).Error
	if err == nil {
		return responseError.ReturnWithOuterError(c, "user already existed")
	}

	err = g.DB.Create(&model.User{
		Username:  username,
		Password:  password,
		Authority: model.AUser,
	}).Error
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	return responseOK.ReturnWithSimpleMessage(c, "log on success")
}

// ModifyUser 修改用户信息 PUT
// 传入格式:
// targetID: 被修改用户id
// authority: 被修改用户新权限
func ModifyUser(c *fiber.Ctx) error {
	type ModifyUserInput struct {
		TargetID  uint32
		Authority model.Authority
	}

	operatorID, err := utils.GetInformationFromJWTToken(c, "id")
	if err != nil {
		return responseError.ReturnWithNotLogin(c)
	}

	user := new(model.User)
	err = g.DB.First(user, operatorID).Error
	if err == gorm.ErrRecordNotFound {
		return responseError.ReturnWithNoUserFound(c)
	} else if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}
	if user.Authority > model.AAdministrator {
		return responseError.ReturnWithNotAuthorize(c, uint32(operatorID.(float64)), operationCode.ModifyUserAuthorize)
	}

	in := new(ModifyUserInput)
	err = c.BodyParser(in)
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	targetID := in.TargetID
	authority := in.Authority
	if targetID == 0 || authority == 0 {
		return responseError.ReturnWithInvalidInput(c)
	}

	targetUser := new(model.User)
	err = g.DB.First(targetUser, targetID).Error
	if err != nil {
		return responseError.ReturnWithNoUserFound(c)
	}

	err = g.DB.Model(targetUser).Update("authority", authority).Error
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	return responseOK.ReturnWithSimpleMessage(c, "modified success")
}

// Login 登录并返回jwt令牌 POST
// 传入信息:
// username: 用户名
// password: 密码
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	in := new(LoginInput)
	err := c.BodyParser(in)
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	username := in.Username
	user := new(model.User)
	err = g.DB.First(user, "username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return responseError.ReturnWithNoUserFound(c)
	} else if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	password := in.Password
	truepass := user.Password
	if password != truepass {
		return responseError.ReturnWithWrongPassword(c)
	}

	expireTime := time.Now().Add(time.Hour * 72)
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": expireTime.Unix(),
	}
	token, err := utils.CreateJWTToken(&claims)
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "session"
	cookie.Value = token
	cookie.Expires = expireTime
	c.Cookie(cookie)
	return responseOK.ReturnWithSimpleMessage(c, "log in success")
}

// GetRestricted 获取限制级页面 GET
// 仅做测试用
func GetRestricted(c *fiber.Ctx) error {
	id, err := utils.GetInformationFromJWTToken(c, "id")
	if err != nil {
		return responseError.ReturnWithNotLogin(c)
	}

	user := new(model.User)
	err = g.DB.First(user, id).Error

	if err == gorm.ErrRecordNotFound {
		return responseError.ReturnWithNoUserFound(c)
	}
	return responseOK.ReturnJsonWithOK(c, fiber.Map{"username": user.Username})
}
