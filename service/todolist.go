package service

import (
	"fmt"

	"github.com/YanJieMao/ToDo/database"
	"github.com/YanJieMao/ToDo/model/pojo"
)

//ToDoListService
type ToDoListService struct {
}

// QueryByUID
func (todolistService *ToDoListService) Query(uid int64) ([]pojo.ToDoList, error) {

	var todoList []pojo.ToDoList
	fmt.Println("tmpDB出错了")

	//tmpDB := todolistService.db.Where("uid = ?", uid)
	tmpDB := database.DB.Where("uid = ?", uid)

	if err := tmpDB.Find(&todoList); err != nil {

		return nil, err
	}

	return todoList, nil

}

//Insert a new todolist and return id
func (todolistService *ToDoListService) Insert(todolist pojo.ToDoList) (int64, error) {
	if _, err := database.DB.Insert(&todolist); err != nil {
		return 0, err
	}

	return todolist.ID, nil

}

//Update  todolist and return id
func (todolistService *ToDoListService) Update(todolist pojo.ToDoList) (int64, error) {
	if _, err := database.DB.Update(&todolist); err != nil {
		return 0, err
	}

	return todolist.ID, nil

}
