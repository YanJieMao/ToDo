package service

import "ToDo/todo/database"

// NewUser get a user service
func NewUser() UserService {
	return UserService{
		db: database.DB,
	}
}
