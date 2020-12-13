package service

import "github.com/YanJieMao/ToDo/todo/database"

// NewUser get a user service
func NewUser() UserService {
	return UserService{
		db: database.DB,
	}
}
