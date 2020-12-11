package service

import (
	"fmt"
	"todo-mvc/todo/datamodles"
	"todo-mvc/todo/datasource"

	"xorm.io/xorm"
)

/**
 * 用户服务实现结构体
 */
type userSevice struct {
	engine *xorm.Engine
}

//定义UserService接口
type UserService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByUserNameAndPassword(username, password string) (datamodles.User, bool)
}

func NewUserService(db *xorm.Engine) UserService {

	return &userSevice{
		engine: db,
	}
}

/**
 * 通过用户名和密码查询管理员
 */
func (ac *userSevice) GetByUserNameAndPassword(username, password string) (datamodles.User, bool) {
	engine := datasource.NewMysqlEngine() //暂时把engine放置在这里
	var user datamodles.User

	engine.Where("username = ? and password = ? ", username, password).Get(&user)

	fmt.Println(user, "............", user.UserId != 0)

	return user, user.UserId != 0
}
