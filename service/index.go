package service

// NewUser get a user service
func NewUser() UserService {

	return UserService{
		//db: database.DB,
	}
}

// NewToDolist get a  todolist service
func NewToDoList() ToDoListService {

	return ToDoListService{
		//db: database.DB,
	}
}
