package service

import "github.com/YanJieMao/ToDo/todo/database"

// NewMessage get a message service
func NewMessage() MessageService {
	return MessageService{
		db: database.DB,
	}
}

// NewUser get a user service
func NewUser() UserService {
	return UserService{
		db: database.DB,
	}
}
