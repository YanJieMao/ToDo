package route

import (
	"github.com/YanJieMao/ToDo/controller"
	"github.com/YanJieMao/ToDo/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeToken(party router.Party) {
	party.Get("/token/info", middleware.JWT.Serve, middleware.Logined, controller.GetTokenInfo)
}
