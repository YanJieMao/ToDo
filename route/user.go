package route

import (
	"github.com/YanJieMao/ToDo/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeUser(party router.Party) {
	party.Post("/login", controller.PostLogin)

	party.Post("/user", controller.PostUser)
	party.Get("/user", controller.GetUser)
	/* party.Put("/user", middleware.JWT.Serve, middleware.Logined, controller.PutUser) */
	party.Put("/user", controller.PutUser)
}
