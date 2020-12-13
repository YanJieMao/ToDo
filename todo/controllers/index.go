package controllers

import (
	"ToDo/todo/database"
	"ToDo/todo/service"
)

var db = database.DB

var userService = service.NewUser()
