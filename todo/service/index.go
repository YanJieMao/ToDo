package service

import "github.com/JabinGP/demo-chatroom/database"

// NewUser get a user service
func NewUser() UserService {
	return UserService{
		db: database.DB,
	}
}
