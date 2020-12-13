package controllers

import (
<<<<<<< HEAD
	"ToDo/todo/datamodels/reso"

	"github.com/JabinGP/demo-chatroom/model"
=======
	
	"github.com/YanJieMao/ToDo/todo/datamodels/reso"
>>>>>>> 0e15918486871b90d2b725f099b6fb6b3df95c2d
	"github.com/kataras/iris/v12"
)

// GetTokenInfo 验证token是否有效，如果有效则返回token携带的信息
func GetTokenInfo(ctx iris.Context) {
	logined := ctx.Values().Get("logined").(model.Logined)

	res := reso.GetTokenInfo{
		ID:       logined.ID,
		Username: logined.Username,
	}
	ctx.JSON(res)
}
