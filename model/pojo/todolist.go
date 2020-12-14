package pojo

import "time"

// TODOList实体，对应表todolist
type ToDoList struct {
	ID        int64
	UID       int64  //用户id
	List      string //内容
	IsDone    int    //0未完成 1完成
	DeadLine  time.Time
	CreatedAt time.Time `xorm:"created"` // 这个Field将在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"updated"` // 这个Field将在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"deleted"` // 如果带DeletedAt这个字段和标签，xorm删除时自动软删除
}
