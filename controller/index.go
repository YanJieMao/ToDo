package controller

import (
	"github.com/YanJieMao/ToDo/service"
)

//var db = database.DB
var userService = service.NewUser()
var todolistService = service.NewToDoList()
