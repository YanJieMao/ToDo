package controller

import (
	"github.com/YanJieMao/ToDo/database"
	"github.com/YanJieMao/ToDo/service"
)

var db = database.DB
var messageService = service.NewMessage()
var userService = service.NewUser()
var todolistService = service.NewToDoList()
