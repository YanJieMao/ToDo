package route

import (
	"github.com/YanJieMao/ToDo/controller"
	"github.com/YanJieMao/ToDo/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeMessage(party router.Party) {
	party.Post("/message", middleware.JWT.Serve, middleware.Logined, controller.PostMessage)
	party.Get("/message", middleware.JWT.Serve, middleware.Logined, controller.GetMessage)
}
