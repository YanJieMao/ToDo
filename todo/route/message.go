package route

import (
	"github.com/YanJieMao/ToDo/todo/controllers"
	"github.com/YanJieMao/ToDo/todo/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeMessage(party router.Party) {
	party.Post("/message", middleware.JWT.Serve, middleware.Logined, controllers.PostMessage)
	party.Get("/message", middleware.JWT.Serve, middleware.Logined, controllers.GetMessage)
}
