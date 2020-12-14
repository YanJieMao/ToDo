package controller

import (
	"github.com/YanJieMao/ToDo/model"
	"github.com/YanJieMao/ToDo/model/reqo"
	"github.com/YanJieMao/ToDo/model/reso"
	"github.com/kataras/iris/v12"
)

func PostToDoList(ctx iris.Context) {
	req := reqo.PostToDoList{}
	ctx.ReadJSON(&req)

}

// GetToDoList return todolist
func GetToDOList(ctx iris.Context) {
	req := reqo.GetToDoList{}
	ctx.ReadJSON(&req)
	resList := []reso.GetToDoList{}

	todoList, err := todolistService.Query(req.UID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	for _, todo := range todoList {
		res := reso.GetToDoList{
			ID:       todo.ID,
			UID:      todo.UID,
			List:     todo.List,
			IsDone:   todo.IsDone,
			DeadLine: todo.DeadLine,
		}

		resList = append(resList, res)
	}

	ctx.JSON(resList)

}
