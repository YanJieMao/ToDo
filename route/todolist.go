package route

import (
	"github.com/YanJieMao/ToDo/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeToDOList(party router.Party) {
	party.Post("/todolist", controller.PostToDoList)
	party.Get("/todolist", controller.GetToDOList)

}
