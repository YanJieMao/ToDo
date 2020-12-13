package route

import (
	"github.com/YanJieMao/ToDo/todo/controllers"
	"github.com/YanJieMao/ToDo/todo/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeToken(party router.Party) {
	party.Get("/token/info", middleware.JWT.Serve, middleware.Logined, controllers.GetTokenInfo)
}
