package service

import (
	"github.com/YanJieMao/ToDo/database"
)

// NewUser get a user service
func NewUser() UserService {
	return UserService{
		db: database.DB,
	}
}

func NewToDoList() ToDoListService {
	return ToDoListService{
		db: database.DB,
	}
}
