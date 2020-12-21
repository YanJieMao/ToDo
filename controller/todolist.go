package controller

import (
	"errors"
	"fmt"

	"github.com/YanJieMao/ToDo/model"
	"github.com/YanJieMao/ToDo/model/pojo"
	"github.com/YanJieMao/ToDo/model/reqo"
	"github.com/YanJieMao/ToDo/model/reso"
	"github.com/kataras/iris/v12"
)

func PostToDoList(ctx iris.Context) {
	req := reqo.PostToDoList{}
	ctx.ReadJSON(&req)
	fmt.Println(req)
	fmt.Println(ctx)

	if req.UID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(errors.New("用户id不能为空")))
		return
	}

	newTodo := pojo.ToDoList{
		UID:      req.UID,
		List:     req.List,
		IsDone:   req.IsDone,
		DeadLine: req.DeadLine,
	}

	todoID, err := todolistService.Insert(newTodo)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorInsertDatabase(err))
		return

	}

	res := reso.PostToDoList{
		ID: todoID,
	}

	ctx.JSON(res)

}

// GetToDoList return todolist
func ToDoListAll(ctx iris.Context) {
	fmt.Println("getToDoList启动了")
	req := reqo.GetToDoList{}
	fmt.Println(ctx)
	ctx.ReadJSON(&req)
	fmt.Println(req)
	resList := []reso.GetToDoList{}

	fmt.Println(resList)

	todoList, err := todolistService.Query(req.UID)
	fmt.Println(todoList)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		fmt.Print("controller查询出错了")
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
