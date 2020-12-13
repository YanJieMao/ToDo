package controllers

import (
	"github.com/YanJieMao/ToDo/todo/database"
	"github.com/YanJieMao/ToDo/todo/service"
)

var db = database.DB

var userService = service.NewUser()
