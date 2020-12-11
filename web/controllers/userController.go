package controllers

import (
	"encoding/json"
	"todo-mvc/todo/service"

	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

/**
 * 用户控制器
 */
type UserController struct {
	//iris框架自动为每个请求都绑定上下文对象：可作为接受参数
	Ctx iris.Context

	//User功能实体：引入Service接口
	Service service.UserService

	//session对象：存储session信息
	Session *sessions.Session
}

const (
	USER = "user" //用户登录成功后存储的session信息的key
)

//将发送请求的字段映射为指定字段
type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

/**
 * 登录功能：json请求格式
 * 接口：/user/login
 */
func (ac *UserController) PostLogin(context iris.Context) mvc.Result {

	var userLogin UserLogin
	ac.Ctx.ReadJSON(&userLogin) //自动将请求的json字符串映射为Login结构体

	//数据参数检验
	if userLogin.UserName == "" || userLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码为空,请重新填写后尝试登录",
			},
		}
	}

	//根据用户名、密码到数据库中查询对应的管理信息
	user, exist := ac.Service.GetByUserNameAndPassword(userLogin.UserName, userLogin.Password)

	//用户不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"success": "登录失败",
				"message": "用户名或者密码错误,请重新登录",
			},
		}
	}

	//用户存在 设置session
	userByte, _ := json.Marshal(user)
	ac.Session.Set(USER, userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登录成功",
			"message": "用户登录成功",
		},
	}
}
