package datamodles

type User struct {
	//如果field名称为Id，而且类型为int64，并没有定义tag，则会被xorm视为主键，并且拥有自增属性
	UserId   int    `xorm:"pk autoincr 'uid'"  json:"uid"` // 主键 自增
	UserName string `xorm:"varchar(32)" json:"username"`
	PassWord string `xorm:"varchar(255)" json:"password"`
}
