package v1

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"b.carriage.fun/datamodel"
	g "b.carriage.fun/server/global"
	responseError "b.carriage.fun/server/response/error"
	responseOK "b.carriage.fun/server/response/ok"
	"b.carriage.fun/server/utils"
)

// AddBangumi 上传bangumi信息 POST
// 传入信息:
// bangumiName: bangumi 名
// webUrl: bangumi 链接
// picUrl: 图片链接
func AddBangumi(c *fiber.Ctx) error {
	type AddBangumiInput struct {
		Name   string
		WebUrl string
		PicUrl string
	}

	in := new(AddBangumiInput)
	err := c.BodyParser(in)
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	userId, err := utils.GetInformationFromJWTToken(c, "id")
	if err != nil {
		return responseError.ReturnWithNotLogin(c)
	}
	user := new(datamodel.User)
	err = g.DB.First(user, userId).Error
	if err == gorm.ErrRecordNotFound {
		return responseError.ReturnWithNoUserFound(c)
	}

	var rs []datamodel.BangumiItem
	err = g.DB.Find(&rs, "name = ?", in.Name).Error
	if err != nil {
		return responseError.ReturnWithOuterError(c, "bangumi already existed")
	}
	bangumiItem := &datamodel.BangumiItem{
		CretedBy: *user,
		Name:     in.Name,
		WebUrl:   in.WebUrl,
		PicUrl:   in.PicUrl,
	}
	err = g.DB.Create(bangumiItem).Error
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	return responseOK.ReturnWithSimpleMessage(c, "新增成功")
}

func GetAllBangumi(c *fiber.Ctx) error {
	var rs = make([]datamodel.BangumiItem, 0, 20)
	err := g.DB.Limit(20).Select([]string{"name", "web_url", "pic_url"}).Find(&rs).Error
	if err != nil {
		return responseError.ReturnWithInternalError(c, err)
	}

	type ReturnedData struct {
		Name   string
		WebUrl string
		PicUrl string
	}
	var rd = make([]ReturnedData, len(rs))
	for i := range rs {
		rd[i].Name = rs[i].Name
		rd[i].WebUrl = rs[i].WebUrl
		rd[i].PicUrl = rs[i].PicUrl
	}

	return responseOK.ReturnJsonWithOK(c, fiber.Map{
		"data": rd,
	})
}
