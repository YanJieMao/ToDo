package route

import (
	"github.com/YanJieMao/ToDo/todo/controllers"
	"github.com/YanJieMao/ToDo/todo/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeUser(party router.Party) {
	party.Post("/login", controllers.PostLogin)

	party.Post("/user", controllers.PostUser)
	party.Get("/user", controllers.GetUser)
	party.Put("/user", middleware.JWT.Serve, middleware.Logined, controllers.PutUser)
}
