package service

import (
	"github.com/YanJieMao/ToDo/model/pojo"
	"xorm.io/xorm"
)

//ToDoListService
type ToDoListService struct {
	db *xorm.Engine
}

// QueryByUID
func (todolistService *ToDoListService) Query(uid int64) ([]pojo.ToDoList, error) {

	var todoList []pojo.ToDoList

	tmpDB := todolistService.db.Where("uid = ?", uid)

	if err := tmpDB.Find(&todoList); err != nil {
		return nil, err
	}

	return todoList, nil

}

//Insert a new todolist and return id
func (todolistService *ToDoListService) Insert(todolist pojo.ToDoList) (int64, error) {
	if _, err := todolistService.db.Insert(&todolist); err != nil {
		return 0, err
	}

	return todolist.ID, nil

}

//Update a new todolist and return id
func (todolistService *ToDoListService) Update(todolist pojo.ToDoList) (int64, error) {
	if _, err := todolistService.db.Update(&todolist); err != nil {
		return 0, err
	}

	return todolist.ID, nil

}
