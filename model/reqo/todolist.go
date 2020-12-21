package reqo

//GEt  /todolist/ request object
type GetToDoList struct {
	ID       int64
	UID      int64  //用户id
	List     string //内容
	IsDone   int    //0未完成 1完成
	DeadLine string
}

// Post  /todolist/ request object
type PostToDoList struct {
	UID      int64  //用户id
	List     string //内容
	IsDone   int    //0未完成 1完成
	DeadLine string
}
